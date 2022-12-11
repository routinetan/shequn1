package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/json-iterator/go/extra"
	"shequn1/internal/biz"
	"shequn1/internal/foundation/view"
)

//管理网站开关,管理网站的标题等
//多个地区的微信号
//多个地区的微信群
type Member struct {
}

func (member Member) List(ctx *gin.Context) {
	num := ctx.DefaultQuery("num", "0")
	ret := biz.GetOpGroupList(gconv.Int(num))
	ret["code"] = 200
	ret["msg"] = ""
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, ret)
}

func (member Member) Info(ctx *gin.Context) {
	view.View.AddPath("group")
	data, _ := view.View.Parse(ctx, "info.tmpl")
	ctx.Status(200)
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.Writer.WriteString(data)
}
