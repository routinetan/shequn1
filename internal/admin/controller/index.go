package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"shequn1/foundation/app"
	"shequn1/foundation/view"
)

// List 自定义 List 方法
func Index(ctx *gin.Context) {
	app.Logger().Println("called this method")
	view.View.AddPath("/dist")
	fp, _ := os.Open("./web/admin/dist/index.html")
	ctx.Status(200)
	io.Copy(ctx.Writer, fp)
}
