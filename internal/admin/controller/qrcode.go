package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/json-iterator/go/extra"
	"shequn1/internal/biz"
	"shequn1/internal/entities"
)

type Qrcode struct {
}

func (qrcode Qrcode) List(ctx *gin.Context) {
	ret := g.Map{}
	groupId := gconv.Int(ctx.Query("group_id"))
	page := gconv.Int(ctx.Query("page"))
	if groupId == 0 {
		ret["code"] = 400
		ret["msg"] = "群id参数有误(请检查群是否正常开启或存在)"
		ctx.PureJSON(400, ret)
		return
	}
	ret = biz.GetQrcodeList(page, groupId)
	ret["code"] = 200
	ret["msg"] = ""
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, ret)
}

func (qrcode Qrcode) Create(ctx *gin.Context) {

	form := entities.QrcodeForm{}
	ret := g.Map{}
	if err := ctx.ShouldBind(&form); err != nil {
		extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
		ret["code"] = 400
		ret["msg"] = err.Error()
		ctx.PureJSON(200, ret)
		return
	}
	ret = biz.CreateQrcode(form)
	ctx.PureJSON(200, ret)
	return
}
