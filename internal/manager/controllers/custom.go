package controllers

import (
	"github.com/gin-gonic/gin"
	"quanzi1/foundation/app"
	"quanzi1/foundation/database/managers"
)

// CustomOrder 继承 mangers 的 MysqlManager 并实现自定义的 List 方法
type CustomOrder struct {
	managers.GormManager
}

// List 自定义 List 方法
func (custom *CustomOrder) List(ctx *gin.Context) {
	app.Logger().Println("called this method")
	custom.GormManager.List(ctx)
}
