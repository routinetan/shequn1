package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"quanzi1/foundation/app"
	"quanzi1/foundation/view"
	"quanzi1/internal/service"
)

// List 自定义 List 方法
func Index(ctx *gin.Context) {
	app.Logger().Println("called this method")
	view.View.AddPath("index")
	view.View.BindFunc("urlfor", func(name string) string {
		var url string
		if name == "/" {
			url = fmt.Sprintf("schema://%s", ctx.Request.Host)
		} else {
			url = fmt.Sprintf("schema://%s/%s", ctx.Request.Host, name)
		}
		return url
	})
	navs := service.GetNavList()
	data, _ := view.View.Parse(ctx, "index.tmpl", g.Map{"navs": navs})
	ctx.Status(200)
	ctx.Writer.WriteString(data)
}
