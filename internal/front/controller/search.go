package controller

import (
	"github.com/gin-gonic/gin"
	"shequn1/internal/biz"
	"shequn1/internal/foundation/view"
)

type Search struct {
}

func (search Search) Index(ctx *gin.Context) {
	content := ctx.PostForm("searchContent")
	view.View.AddPath("search")
	ret := biz.SearchGroupList(content)
	data, _ := view.View.Parse(ctx, "search.tmpl", ret)
	ctx.Status(200)
	ctx.Writer.WriteString(data)
}
