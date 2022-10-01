package doc

import (
	"github.com/gin-gonic/gin"
	"quanzi1/foundation/middlewares"
	"quanzi1/internal/doc/controllers/order"
)

// GetEngine 路由注册主方法
// @title starter
// @version 1.0
// @host golang-project.com
func GetEngine(engine *gin.Engine) {
	engine.Use(middlewares.CORS)
	engine.GET("/order", order.List)

	var category = order.NewCategory()
	categories := engine.Group("/categories")
	categories.GET("/mgo", category.Mgo)
	categories.GET("/mongo", category.Mongo)
	categories.GET("/mysql", category.Mysql)
	categories.GET("/list_mgo", category.ListMgo)
	categories.GET("/list_mysql", category.ListMysql)
	categories.GET("/list_mongo", category.ListMongo)
}
