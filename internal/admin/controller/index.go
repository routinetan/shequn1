package controller

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"shequn1/internal/foundation/app"
	"shequn1/internal/foundation/view"
)

// List 自定义 List 方法
func Index(ctx *gin.Context) {
	app.Logger().Println("called this method")
	view.View.AddPath("/dist")
	fp, _ := os.Open("./view/admin/dist/index.html")
	ctx.Status(200)
	io.Copy(ctx.Writer, fp)
}
