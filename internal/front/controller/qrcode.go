package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/util/gconv"
	"io"
	"os"
	"shequn1/internal/biz"
	"shequn1/internal/foundation/view"
)

type Qrcode struct {
}

func (qrcode Qrcode) Index(ctx *gin.Context) {

	groupIdstr := ctx.Param("group_id")
	groupId := gconv.Int(groupIdstr)

	view.View.AddPath("qrcode")
	//html, _ := view.View.Parse(ctx, "index.tmpl")

	qrcodeThumb := biz.UserScanLogic(groupId)
	if qrcodeThumb != "" {
		f, _ := os.Open(qrcodeThumb)
		io.Copy(ctx.Writer, f)
		return
	}
	//
	//ctx.Redirect(308, url)

}
