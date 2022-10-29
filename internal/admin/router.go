package admin

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"html/template"
	"net/http"
	"os"
	"shequn1/foundation/app"
	"shequn1/foundation/database/managers"
	"shequn1/foundation/middlewares"
	"shequn1/foundation/rbac"
	"shequn1/foundation/util"
	"shequn1/internal/admin/controller"
	"shequn1/internal/entities"
	"shequn1/internal/manager/controllers"
)

func GetRouter(engine *gin.Engine) {
	engine.LoadHTMLFiles()
	// 注册自定义函数
	engine.SetFuncMap(template.FuncMap{
		"map": func(json string) gin.H {
			var out gin.H
			_ = jsoniter.UnmarshalFromString(json, &out)
			return out
		},
	})

	// 加载模板
	//engine.LoadHTMLGlob("web/admin/dist/*")
	engine.StaticFS("/assets", http.Dir("./web/admin/dist/assets/"))
	engine.StaticFS("/css", http.Dir("./web/admin/dist/css/"))
	engine.StaticFS("/js", http.Dir("./web/admin/dist/js/"))
	engine.StaticFS("/public", http.Dir("./public"))
	// 注册公用的中间件
	engine.Use(middlewares.CORS)

	// 登录路由需要在jwt验证中间件之前
	engine.GET("/", controller.Index)

	// 注册一个权限验证的中间件
	//engine.Use(managerMiddleWares.CheckPermission)

	// 注册一个公共上传接口
	if !util.Exists("./public/wx") {
		os.MkdirAll("./public/wx", 0777)
	}
	var saveHandler = new(app.DefaultSaveHandler).SetPrefix("http://localhost:8082")
	saveHandler.SetPublicDst("/public/wx/").SetDst("./public/wx/")
	engine.POST("/upload", app.Upload("file", saveHandler, "png", "jpg"))
	apiv1 := engine.Group("/apiv1")
	apiv1.GET("/groups", controller.Group{}.List)
	apiv1.GET("/group/:id", controller.Group{}.Info)
	apiv1.GET("/qrcodes", controller.Qrcode{}.List)
	apiv1.POST("/qrcode/add", controller.Qrcode{}.Create)
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
