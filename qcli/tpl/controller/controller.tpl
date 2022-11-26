package controller
import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/json-iterator/go/extra"
	"shequn1/foundation/view"
	"shequn1/internal/service"
)

//管理网站开关,管理网站的标题等
type {{.CStruct}} struct {

}

type Edit{{.CStruct}}BizRule struct {

}

type Create{{.CStruct}}BizRule struct {

}

func ({{.Cval}} {{.CStruct}}) List(ctx *gin.Context) {
    resp := g.Map{}
	resp["code"] = 200
	resp["msg"] = ""

	num := ctx.DefaultQuery("num", "0")
	resp['data'] := service.Get{{.CStruct}}List(gconv.Int(num))
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, resp)
}

func ({{.Cval}} {{.CStruct}}) Info(ctx *gin.Context) {
    resp := g.Map{}
	resp["code"] = 200
	resp["msg"] = ""

	id := ctx.DefaultQuery("id", "")
	err,resp['data']  := service.Get{{.CStruct}}(gconv.Int(id))
	 if err != nil {
        resp["code"] = 400
        resp["msg"] = err.Error()
        ctx.PureJSON(200, ret)
        return
    }
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, resp)
}

func ({{.Cval}} {{.CStruct}}) Create(ctx *gin.Context) {
    resp := g.Map{}
    //请求参数校验
    bizRule := Edit{{.CStruct}}BizRule{}
    //参数绑定 修改时候的参数和更新时候的参数。
    bizReq := g.Map{}
    verr := app.ValidatorRules(ctx, bizRule, &bizReq)
    if verr != nil {
        resp["code"] = 400
        resp["msg"] = verr.Error()
        ctx.PureJSON(200, resp)
        return
    }
    err,resp['data'] := service.Create{{.CStruct}}(bizReq)
    if err != nil {
        resp["code"] = 400
        resp["msg"] = err.Error()
        ctx.PureJSON(200, resp)
        return
    }
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, resp)
}

func ({{.Cval}} {{.CStruct}}) Edit(ctx *gin.Context) {
    //参数绑定
    resp := g.Map{}
    bizRule := Edit{{.CStruct}}BizRule{}
    bizReq := g.Map{}
    verr := app.ValidatorRules(ctx, bizRule, &bizReq)
    if verr != nil {
        resp["code"] = 400
        resp["msg"] = verr.Error()
        ctx.PureJSON(200, resp)
        return
    }
    err,resp['data'] := service.Edit{{.CStruct}}(bizReq)
    if err != nil {
        resp["code"] = 400
        resp["msg"] = err.Error()
        ctx.PureJSON(200, resp)
        return
    }
	ctx.Status(200)
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, resp)
}


func ({{.Cval}} {{.CStruct}}) Delete(ctx *gin.Context) {
    resp := g.Map{}
    //请求参数校验
    bizRule := map[string]string{
    }
    //参数绑定 修改时候的参数和更新时候的参数。
    bizReq := g.Map{}
    verr := app.ValidatorRules(ctx, bizRule, &bizReq)
    if verr != nil {
        resp["code"] = 400
        resp["msg"] = verr.Error()
        ctx.PureJSON(200, resp)
        return
    }
    err := service.Delete{{.CStruct}}(bizReq)
    if err != nil {
        resp["code"] = 400
        resp["msg"] = err.Error()
        ctx.PureJSON(200, resp)
        return
    }
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, resp)
}



