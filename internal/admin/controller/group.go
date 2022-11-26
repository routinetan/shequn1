package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/json-iterator/go/extra"
	"shequn1/foundation/app"
	"shequn1/foundation/validator"
	"shequn1/internal/service"
)

//管理网站开关,管理网站的标题等
//多个地区的微信号
//多个地区的微信群
type Group struct {
}

type EditBizRule struct {
}

type CreateGroupBizReq struct {
	Id       int    `controller:"id,omitempty"`
	Title    string `controller:"title,omitempty"`
	Label    string `controller:"label,omitempty"`
	ClassId  int    `controller:"class_id,omitempty"`
	Type     int    `controller:"type,omitempty"`
	ThumbUrl string `controller:"thumb_url,omitempty"`
	IsShow   string `controller:"is_show,omitempty"`
	Status   int    `controller:"status,omitempty"`
}

type EditBizReq struct {
	Id       int    `controller:"id,omitempty"`
	Title    string `controller:"title,omitempty"`
	Label    string `controller:"label,omitempty"`
	ClassId  int    `controller:"class_id,omitempty"`
	ThumbUrl string `controller:"thumb_url,omitempty"`
	IsShow   string `controller:"is_show,omitempty"`
	Status   int    `controller:"status,omitempty"`
}

//rules和req

func (group Group) List(ctx *gin.Context) {
	num := ctx.DefaultQuery("num", "0")
	ret := service.GetOpGroupList(gconv.Int(num))
	ret["code"] = 200
	ret["msg"] = ""
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, ret)
}

func (group Group) Info(ctx *gin.Context) {
	id := ctx.Param("id")
	ret := g.Map{}
	ret["code"] = 200
	ret["msg"] = ""
	ret["data"] = service.GetGroupInfo(gconv.Int(id))
	ctx.Status(200)
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, ret)
}

func (group Group) CreateGroupBase(ctx *gin.Context) {
	ret := g.Map{}
	ret["code"] = 200
	ret["msg"] = ""
	//请求参数校验
	bizRule := map[string]string{
		"title":    "required|min-length:1",
		"label":    "required|min-length:1",
		"is_show":  "integer|min:0",
		"class_id": "required|integer|min:0",
	}
	//参数绑定
	bizReq := CreateGroupBizReq{}
	verr := app.ValidatorRules(ctx, bizRule, &bizReq)
	if verr != nil {
		ret["code"] = 400
		ret["msg"] = verr.Error()
		ctx.PureJSON(200, ret)
		return
	}
	bizAttrs := gconv.Map(bizReq)
	err, groupBizObj := service.CreateGroupBase(bizAttrs)
	if err != nil {
		ret["code"] = 400
		ret["msg"] = err.Error()
		ctx.PureJSON(200, ret)
		return
	}
	ret["data"] = groupBizObj
	ctx.PureJSON(200, ret)
}

func (group Group) ModifyGroupBase(ctx *gin.Context) {
	id := gconv.Int(ctx.Param("id"))
	ret := g.Map{}
	ret["code"] = 200
	ret["msg"] = ""
	//请求参数校验
	bizRule := map[string]string{
		"id":      "integer",
		"cate_id": "integer",
		"status":  "integer",
	}
	//参数绑定
	bizReq := EditBizReq{}
	verr := app.ValidatorRules(ctx, bizRule, &bizReq)
	if verr != nil {
		ret["code"] = 400
		ret["msg"] = verr.Error()
		ctx.PureJSON(200, ret)
		return
	}
	bizAttrs := gconv.Map(bizReq)
	err := service.ModifyGroupBase(id, bizAttrs)
	if err != nil {
		ret["code"] = 400
		ret["msg"] = err.Error()
		ctx.PureJSON(200, ret)
		return
	}
	ctx.PureJSON(200, ret)
}

func (group Group) ModifyGroupAdvance(ctx *gin.Context) {
	ret := g.Map{}
	ret["code"] = 200
	ret["msg"] = ""
	var groupBaseForm EditBizReq

	if err := validator.Bind(ctx, &groupBaseForm); !err.IsValid() {
		app.NewResponse(app.Success, err.ErrorsInfo).End(ctx)
		return
	}
	groupEntity := g.Map{}
	gconv.Structs(groupBaseForm, &groupEntity)
	err := service.ModifyGroupBase(groupBaseForm.Id, groupEntity)
	if err != nil {
		ret["code"] = 400
		ret["msg"] = err.Error()
	}
	ctx.PureJSON(200, ret)
}
