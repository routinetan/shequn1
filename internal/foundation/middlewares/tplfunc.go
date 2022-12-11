package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"shequn1/internal/foundation/view"
)

func RegisterFuc(ctx *gin.Context) {
	view.View.BindFunc("urlfor", func(name string) string {
		var url string
		if name == "/" {
			url = fmt.Sprintf("https://%s", ctx.Request.Host)
		} else {
			url = fmt.Sprintf("https://%s/%s", ctx.Request.Host, name)
		}
		return url
	})
}
