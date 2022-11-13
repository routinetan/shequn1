package front

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
	"shequn1/foundation/view"
	"shequn1/internal/entities"
	"shequn1/internal/front/controller"
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
	view.View.AddPath("front/tmpl")
	// 加载模板
	engine.LoadHTMLGlob("web/front/tmpl/*/*")
	engine.StaticFS("/static", http.Dir("./web/static"))
	engine.StaticFS("/public", http.Dir("./public"))
	// 注册公用的中间件
	engine.Use(middlewares.CORS)
	engine.Use(middlewares.RegisterFuc)
	// 登录路由需要在jwt验证中间件之前
	engine.GET("/", controller.Index)
	engine.GET("/group_list", controller.Group{}.List)
	engine.GET("/group/info", controller.Group{}.Info)

	engine.POST("/search", controller.Search{}.Index)
	engine.GET("/qrcode/:group_id", controller.Qrcode{}.Index)
	// 注册一个权限验证的中间件
	//engine.Use(managerMiddleWares.CheckPermission)
	if !util.Exists("./public/wx") {
		os.MkdirAll("./public/wx", 0777)
	}
	// 注册一个公共上传接口
	var saveHandler = new(app.DefaultSaveHandler).SetDst("./public/wx/")
	engine.POST("/upload", app.Upload("file", saveHandler, "png", "jpg"))

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
