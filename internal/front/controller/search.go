package controller

import (
	"github.com/gin-gonic/gin"
	"shequn1/foundation/view"
	"shequn1/internal/service"
)

type Search struct {
}

func (search Search) Index(ctx *gin.Context) {
	content := ctx.PostForm("searchContent")
	view.View.AddPath("search")
	ret := service.SearchGroupList(content)
	data, _ := view.View.Parse(ctx, "search.tmpl", ret)
	ctx.Status(200)
	ctx.Writer.WriteString(data)
}
