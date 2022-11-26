package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/json-iterator/go/extra"
	"shequn1/foundation/app"
	"shequn1/internal/service"
)

type Setting struct {
}

type SystemCfgModifyBizReq struct {
	Id                     int    `controller:"id"`
	DefaultWechatStatus    int    `controller:"default_wechat_status"`
	DefaultWechatQrCodeUrl string `controller:"default_wechat_qrcodeurl"`
	VipStatus              int    `controller:"vip_status"`
	PayCloseStatus         int    `controller:"pay_close_status"`
	FeederDomainStatus     int    `controller:"feeder_domain_status"`
	GenGroupStatus         int    `controller:"gen_group_status"`
	GenGroupId             int    `controller:"gen_group_id"`
	CloseStatus            int    `controller:"close_status"`
}

func (setting Setting) GetSystemConfig(ctx *gin.Context) {
	ret := g.Map{}
	ret["code"] = 200
	ret["msg"] = ""
	ret["data"] = service.GetSystemCfg()
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, ret)
}

func (setting Setting) SaveSystemConfig(ctx *gin.Context) {
	ret := g.Map{}
	ret["code"] = 200
	ret["msg"] = ""
	//请求参数校验
	bizRule := map[string]string{
		"default_wechat_status":    "integer|min:0",
		"default_wechat_qrcodeurl": "min-length:1|required-if:default_wechat_status,1",
		"vip_status":               "integer|min:0",
		"gen_group_status":         "integer|min:0",
		"gen_group_id":             "integer|required-if:gen_group_status,1",
		"close_status":             "integer|min:0",
		"feeder_domain_status":     "integer|min:0",
		"pay_close_status":         "integer|min:0",
	}
	//参数绑定
	bizReq := g.Map{}
	//根据当前请求的键名,生成对象。如果对象不是map，则返回map
	//绑定请求，过滤请求参数
	verr := app.ValidatorRules(ctx, bizRule, &bizReq)
	if verr != nil {
		ret["code"] = 400
		ret["msg"] = verr.Error()
		ctx.PureJSON(200, ret)
		return
	}
	err := service.SaveSystemCfg(bizReq)
	if err != nil {
		ret["code"] = 400
		ret["msg"] = err.Error()
		ctx.PureJSON(200, ret)
		return
	}
	ctx.PureJSON(200, ret)
}
