//go:build windows
// +build windows

package server

import (
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"shequn1/foundation/app"
	"shequn1/foundation/database/orm"
	"shequn1/foundation/engine"
	"shequn1/foundation/log"
	"shequn1/foundation/redis"
	"shequn1/foundation/validator"
	"shequn1/foundation/view"
	"time"
)

type (
	application struct {
		Name          string `toml:"name"`
		Domain        string `toml:"domain"`
		Addr          string `toml:"addr"`
		PasswordToken string `toml:"password_token"`
		JwtToken      string `toml:"jwt-token"`
		CertFile      string `toml:"cert_file"`
		KeyFile       string `toml:"key_file"`
	}
)

var (
	pidFile = fmt.Sprintf("./%s.pid", app.Name())
	// Mode 当前服务名, 暂时需要这么个东西, 按目前的结构获得指定配置信息等操作
	Mode string
	// After 在各项服务启动之后会执行的操作
	After       func(engine *gin.Engine)
	swagHandler gin.HandlerFunc
	router      = gin.New()

	// Config 服务配置
	Config application
)

func certInfo() (string, string) {
	return Config.CertFile, Config.KeyFile
}

// 启动各项服务
func start() {
	log.Start()
	orm.Start()
	//mongo.Start()
	//mgo.Start()
	redis.Start()
	//elastic.Start()
	view.Init()
	view.View.AddPath("/" + Mode + "/")
	// 加载应用配置
	err := app.Config().Bind("application", fmt.Sprintf("application.%s", Mode), &Config)
	if err != nil {
		fmt.Println(err)
	}
	// 将 gin 的验证器替换为 v9 版本
	binding.Validator = new(validator.Validator)
}

//Run 启动服务
func Run(router func(engine engine.Engine)) {
	//lock := createPid()
	//defer lock.UnLock()

	start()
	app.Logger().WithField("log_type", "foundation.server.server").Info("server started at:", time.Now().String())
	engineInst := engine.GetEngine()
	router(engineInst)

	//if swagHandler != nil && gin.Mode() != gin.ReleaseMode {
	//	engine.GET("/doc/*any", swagHandler)
	//}
	createServer(engineInst.GetHandler()).ListenAndServe()
	//_ = gracehttp.ServeWithOptions([]*http.Server{createServer(engine)}, gracehttp.PreStartProcess(func() error {
	//	app.Logger().WithField("log_type", "foundation.server.server").Println("unlock pid")
	//	lock.UnLock()
	//	return nil
	//}))
}

func createServer(router http.Handler) *http.Server {
	server := &http.Server{
		Addr:    Config.Addr,
		Handler: router,
	}

	if certFile, certKey := certInfo(); certFile != "" && certKey != "" {
		server.TLSConfig = &tls.Config{}
		f, _ := tls.LoadX509KeyPair(certFile, certKey)
		server.TLSConfig.Certificates = []tls.Certificate{f}
	}

	return server
}

// 对启动进程记录进程id
//func createPid() *app.Flock {
//pidLock, pidLockErr := app.FLock(pidFile)
//if pidLockErr != nil {
//	app.Logger().WithField("log_type", "foundation.server.server").Fatalln("createPid: lock error", pidLockErr)
//}

//err := pidLock.WriteTo(fmt.Sprintf(`%d`, os.Getpid()))
//if err != nil {
//	app.Logger().WithField("log_type", "foundation.server.server").Fatalln("write error: ", err)
//}
//return nil
//}
