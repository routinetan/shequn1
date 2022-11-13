package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"shequn1/foundation/app"
	"shequn1/foundation/view"
	"shequn1/internal/service"
)

// List 自定义 List 方法
func Index(ctx *gin.Context) {
	cateId := ctx.DefaultQuery("cate_id", "")
	app.Logger().Println("called this method")
	view.View.AddPath("index")
	navs := service.GetNavList()
	data, _ := view.View.Parse(ctx, "index.tmpl", g.Map{"navs": navs, "CateId": cateId})
	ctx.Status(200)
	ctx.Writer.WriteString(data)
}
