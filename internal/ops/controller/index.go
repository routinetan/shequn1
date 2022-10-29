package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"shequn1/foundation/app"
	"shequn1/foundation/view"
)

// List 自定义 List 方法
func Index(ctx *gin.Context) {
	app.Logger().Println("called this method")
	view.View.AddPath("/index")
	data, _ := view.View.Parse(context.TODO(), "index.tmpl")
	ctx.Status(200)
	ctx.Writer.WriteString(data)
}
