package manager

import (
	"github.com/gin-gonic/gin"
	"shequn1/internal/entities"
	app2 "shequn1/internal/foundation/app"
	"shequn1/internal/foundation/captcha"
	"shequn1/internal/foundation/database/managers"
	"shequn1/internal/foundation/middlewares"
	"shequn1/internal/foundation/rbac"
	"shequn1/internal/manager/controllers"
)

// GetEngine 设置路由的主方法,示例
func GetRouter(engine *gin.Engine) {
	// 静态资源路径, 这里只是临时写了一个文件夹作为示例
	engine.Static("/test", "./test")

	// 注册公用的中间件
	engine.Use(middlewares.CORS)

	// 登录路由需要在jwt验证中间件之前
	engine.POST("/login", controllers.Login)

	engine.GET("/captcha", func(context *gin.Context) {
		cpat := captcha.New("medivh")
		app2.NewResponse(app2.Success, gin.H{"content": cpat.ToBase64EncodeString(), "captcha_id": cpat.CaptchaID}).End(context)
	})

	engine.POST("/captcha", func(context *gin.Context) {
		id := context.DefaultQuery("captcha_id", "medivh")
		app2.Logger().Debug(captcha.Verify(id, context.DefaultQuery("captcha", "")))
	})

	//engine.Use(middlewares.VerifyAuth)
	//sessions.Inject(engine)

	engine.GET("/staffs/info", controllers.StaffInfo)

	// 注册一个权限验证的中间件
	//engine.Use(managerMiddleWares.CheckPermission)

	// 注册一个公共上传接口
	var saveHandler = new(app2.DefaultSaveHandler).SetPrefix("http://manager.golang-project.com/").SetDst("./test/")
	engine.POST("/upload", app2.Upload("file", saveHandler, "png", "jpg"))

	// CSRFtoken支持, 因为 upload 不需要受 CSRFtoken 限制, 故将上传接口放在了上边
	//engine.Use(middlewares.CsrfToken)

	// 将对应数据接口注册生成 CURD 接口
	managers.New().
		Register(entities.Staff{}, managers.Mongo).
		Register(entities.Mgo{}, managers.Mgo).
		RegisterCustomManager(&controllers.CustomOrder{}, entities.Order{}).
		Start(engine)

	// 将权限验证数据表的CURD接口进行注册
	rbac.Inject(engine)
}
