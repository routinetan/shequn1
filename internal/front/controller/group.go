package controller

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/json-iterator/go/extra"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"shequn1/foundation/view"
	"shequn1/internal/service"
)

type Group struct {
}

func (group Group) List(ctx *gin.Context) {
	num := ctx.DefaultQuery("num", "0")
	ret := service.GetGroupList(gconv.Int(num))
	ret["code"] = 200
	ret["msg"] = ""
	extra.SetNamingStrategy(extra.LowerCaseWithUnderscores)
	ctx.PureJSON(200, ret)
}

func (group Group) Info(ctx *gin.Context) {
	view.View.AddPath("group")
	params := g.Map{
		"status": 0,
	}

	table := make(map[int]string)
	table[10] = "197"
	table[20] = "199"
	table[30] = "200"
	table[40] = "198"
	//最后一个为不满群，其他随机，roll一个100内的数, 有记录为 有198 40 199 20 200 30 197 10 三种，最后一个群一定是不满的
	cfg := InitSeedTable(table)
	num := make(map[int]string)
	maxNum := 5
	for i := 1; i < maxNum; i++ {
		num[i] = Poll(cfg)
	}

	num[maxNum] = "198"
	params["Num"] = num
	imgByte, _ := qrcode.Encode(fmt.Sprintf("http://192.168.31.172:8081/qrcode/%d", rand.Int()), qrcode.Medium, 256)
	baseImg := base64.StdEncoding.EncodeToString(imgByte)

	params["QrCode"] = baseImg
	data, _ := view.View.Parse(ctx, "info.tmpl", params)
	ctx.Status(200)
	ctx.Writer.WriteString(data)
}

func InitSeedTable(table map[int]string) map[int]string {
	cfg := make(map[int]string)
	num := 0
	for k, v := range table {
		num = k + num
		cfg[num] = v
	}
	return cfg
}

func Poll(cfg map[int]string) string {
	poll := rand.Int31n(100)
	for k, v := range cfg {
		if int(poll) < k {
			return v
		}
	}
	return ""
}
