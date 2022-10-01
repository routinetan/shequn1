package engine

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"quanzi1/foundation/app"
	"strings"
	"time"
)

type Engine interface {
	Init()
	Run(router func(engine *gin.Engine))
	GetHandler() http.Handler
}

var engine Engine

func GetEngine() Engine {
	return NewGinEngine()
}

type GinEngine struct {
	Gin *gin.Engine
}

func (ginRouter *GinEngine) Init() {
	//ginRouter.Gin.Use(logger, recovery)
}

func (ginRouter *GinEngine) Run(router func(engine *gin.Engine)) {
	router(ginRouter.Gin)
	//ginRouter.Gin.Use(logger, recovery)
}

func (ginRouter *GinEngine) GetHandler() http.Handler {
	return ginRouter.Gin
}

func NewGinEngine() *GinEngine {
	return &GinEngine{Gin: gin.New()}
}

// 自定义的GIN日志处理中间件
// 可能在终端输出时看起来比较的难看
func logger(ctx *gin.Context) {
	start := time.Now()
	path := ctx.Request.URL.Path
	raw := ctx.Request.URL.RawQuery

	ctx.Next()

	if raw != "" {
		path = path + "?" + raw
	}

	var params = make(logrus.Fields)
	params["latency"] = time.Now().Sub(start)
	params["path"] = path
	params["method"] = ctx.Request.Method
	params["status"] = ctx.Writer.Status()
	params["body_size"] = ctx.Writer.Size()
	params["client_ip"] = ctx.ClientIP()
	params["user_agent"] = ctx.Request.UserAgent()
	params["log_type"] = "foundation.server.server"
	if !gin.IsDebugging() {
		// 在正式环境将上下文传递的变量也进行记录, 方便分析
		params["keys"] = ctx.Keys
	}
	app.Logger().WithFields(params).Info("request success, status is ", ctx.Writer.Status(), ", client ip is ", ctx.ClientIP())
}

func recovery(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var brokenPipe bool
			if ne, ok := err.(*net.OpError); ok {
				if se, ok := ne.Err.(*os.SyscallError); ok {
					if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
						brokenPipe = true
					}
				}
			}
			stack := app.Stack(3)
			httpRequest, _ := httputil.DumpRequest(ctx.Request, false)

			if gin.IsDebugging() {
				app.Logger().WithField("log_type", "foundation.server.server").Error(string(httpRequest))
				var errors = make([]logrus.Fields, 0)
				for i := 0; i < len(stack); i++ {
					errors = append(errors, logrus.Fields{
						"func":   stack[i]["func"],
						"source": stack[i]["source"],
						"file":   fmt.Sprintf("%s:%d", stack[i]["file"], stack[i]["line"]),
					})
				}
				app.Logger().WithField("log_type", "foundation.server.server").WithField("stack", errors).Error(err)
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"stack": errors, "message": err})
			} else {
				app.Logger().WithField("log_type", "foundation.server.server").
					WithField("stack", stack).WithField("request", string(httpRequest)).Error()
			}

			if brokenPipe {
				_ = ctx.Error(err.(error))
				ctx.Abort()
			} else {
				ctx.AbortWithStatus(http.StatusInternalServerError)
			}
		}
	}()
	ctx.Next()
}
