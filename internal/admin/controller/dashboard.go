package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/json-iterator/go/extra"
)

type DashBoard struct {
}

func (dashBoard DashBoard) Report(ctx *gin.Context) {
	ret := g.Map{}
	ret["new_user_day"] = 0
	ret["new_group_day"] = 0
	ret["new_order_day"] = 0
	ret["new_recharge_day"] = 0

	ret["user_total"] = 0
	ret["group_total"] = 0
	ret["order_total"] = 0
	ret["recharge_total"] = 0

	ret["code"] = 200
	ret["msg"] = ""
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, ret)
}
