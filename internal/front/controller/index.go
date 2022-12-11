package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	Biz "shequn1/internal/biz"
	"shequn1/internal/foundation/view"
)

// List 自定义 List 方法
func Index(ctx *gin.Context) {
	cateId := ctx.DefaultQuery("cate_id", "")
	view.View.AddPath("index")
	navs := Biz.GetNavList()
	data, _ := view.View.Parse(ctx, "index.tmpl", g.Map{"navs": navs, "CateId": cateId})
	ctx.Status(200)
	ctx.Writer.WriteString(data)
}
