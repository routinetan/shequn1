package service

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/gorm"
	"github.com/tuotoo/qrcode"
	"os"
	"shequn1/foundation/app"
	"shequn1/foundation/database/orm"
	"shequn1/foundation/paginator"
	"shequn1/internal/entities"
	"shequn1/internal/model"
)

type Upload struct {
	Prefix    string
	PublicDst string
	Dst       string
}

func GetQrcodeList(page, groupId int) g.Map {
	var OpQrcodes []entities.OpQrcodeList
	var qrcodes []model.Qrcode
	var countRows int
	orm.Master().Where("group_id = ?", groupId).Find(&qrcodes).Count(&countRows)
	p := paginator.NewPagintor(0, countRows)
	orm.Master().Where("group_id = ?", groupId).Offset(page * p.Rows).Limit(p.Rows).Order("id desc").Find(&qrcodes)
	gconv.Structs(qrcodes, &OpQrcodes)
	rownum := countRows / p.Rows
	var upload Upload
	app.Config().Bind("application", "upload", &upload)
	for k, v := range OpQrcodes {
		OpQrcodes[k].QrcodeUrl = upload.Prefix + v.QrcodeUrl
	}
	return g.Map{"rownum": rownum + 1, "rows": len(OpQrcodes), "data": OpQrcodes}
}

//处理用户扫描二维码的过程
func UserScanLogic(groupId int) string {
	//通过当前群名获取真实二维码
	//判断当前二维码是否有超过最大限制
	//更换二维码,重置当前前台显示入群总人数
	//增加统计人群人数
	//给机器人发日志,通知其开始给进群的用户打招呼
	group := model.Group{Id: groupId}
	orm.Master().First(&group)
	qrcode := model.Qrcode{GroupId: groupId, Status: entities.QRCODE_STATUS_NORMAL}
	orm.Master().First(&qrcode)
	if qrcode.Num > group.MaxNum {
		//换新群
	}
	qrcode.Num = qrcode.Num + 1
	orm.Master().Save(&qrcode)
	return qrcode.QrcodeUrl
}

func CreateQrcode(formData entities.QrcodeForm) g.Map {
	ret := g.Map{}
	ret["error"] = ""
	ret["data"] = ""
	groupId := gconv.Int(formData.GroupId)
	group := model.Group{}
	err := orm.Master().First(&group, "id = ? and status=?", groupId, entities.QRCODE_STATUS_NORMAL).Error
	if err == gorm.ErrRecordNotFound {
		ret["error"] = "群未查到，请检查群已被删除或者非正常状态"
		ret["data"] = ""
		return ret
	}

	qrcode := &model.Qrcode{}
	gconv.Struct(formData, &qrcode)
	createErr := orm.Master().Save(&qrcode).Error
	if createErr != nil {
		app.Logger().Error(err.Error())
		ret["error"] = "添加失败,请稍后再试"
		ret["data"] = ""
		return ret
	}
	return ret
}

func ParserQrcode(imgPath string) string {
	fi, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return qrmatrix.Content

}

//给机器人发消息
func SendRobotMessage() {

}

//造假数据的规则
//最后一个不满
//每个都不满
//其中一边的满
//根据实际的群来设置设置最后一个
//群的头像可以都为女性，或者.或者混合
func BuilderGroupData() {
	//获取生成群

}
