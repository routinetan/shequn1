package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/json-iterator/go/extra"
	"net/http"
	"shequn1/internal/biz"
	"shequn1/internal/foundation/app"
	"shequn1/internal/foundation/view"
)

//管理网站开关,管理网站的标题等
//多个地区的微信号
//多个地区的微信群
type Order struct {
}

func (order Order) List(ctx *gin.Context) {
	resp := app.NewResponse(http.StatusOK, nil, "")
	num := ctx.DefaultQuery("num", "0")
	resp.Data = biz.GetOpGroupList(gconv.Int(num))
	resp.End(ctx)
	return
}

func (order Order) Info(ctx *gin.Context) {
	view.View.AddPath("group")
	data, _ := view.View.Parse(ctx, "info.tmpl")
	ctx.Status(200)
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.Writer.WriteString(data)
}
