package controller

import (
	"github.com/gin-gonic/gin"
	"shequn1/foundation/view"
)

type Order struct {
}

func (order Order) Index(ctx *gin.Context) {
	view.View.AddPath("order")
	data, _ := view.View.Parse(ctx, "order.tmpl")
	ctx.Status(200)
	ctx.Writer.WriteString(data)
}

func (order Order) List() {

}

func (order Order) Submit() {

}

func (order Order) Payment() {

}
