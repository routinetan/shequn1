package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"github.com/json-iterator/go/extra"
	"quanzi1/internal/service"
)

type Group struct {
}

func (group Group) List(ctx *gin.Context) {
	num := ctx.DefaultQuery("num", "8")
	ret := service.GetGroupList(gconv.Int(num))
	ret["code"] = 200
	ret["msg"] = ""
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, ret)
}
