package biz

import (
	"errors"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"github.com/jinzhu/gorm"
	"shequn1/internal/foundation/database/orm"
	"shequn1/internal/foundation/paginator"
	"shequn1/internal/foundation/server"
	"shequn1/internal/store/entities"
	"shequn1/internal/store/model"
)

const (
	_ = iota
	GROUP_SHOW_STATUS_ON
	GROUP_SHOW_STATUS_OFF
)

const (
	_ = iota
	GROUP_STATUS_NEED_PAY
	GROUP_STATUS_FREE
)

func SearchGroupList(content string) g.Map {
	var groups []entities.Group
	var groupList []model.Group
	var countRows int
	orm.Master().Find(&groupList).Count(&countRows)
	orm.Master().Where("is_show = ?", GROUP_SHOW_STATUS_ON).Where("title like ?", "%"+content+"%").Order("rank_score desc").Find(&groupList)
	gconv.Structs(groupList, &groups)
	return g.Map{"Groups": groups}
}

func GetGroupList(num int) g.Map {
	var groups []entities.Group
	var groupList []model.Group
	var countRows int
	orm.Master().Find(&groupList).Count(&countRows)
	p := paginator.NewPagintor(0, countRows)
	orm.Master().Offset(num).Limit(p.Rows).Where("is_show = ?", GROUP_SHOW_STATUS_ON).Order("rank_score desc").Find(&groupList)
	gconv.Structs(groupList, &groups)
	for k, v := range groupList {
		if v.Status == GROUP_STATUS_NEED_PAY {
			groups[k].Jumpbtntext = "收费"
		}
	}
	rownum := countRows / p.Rows
	return g.Map{"rownum": rownum + 1, "rows": len(groups), "data": groups}
}

func GetOpGroupList(num int) g.Map {
	var groups []entities.OpGroup
	var groupList []model.Group
	var countRows int
	orm.Master().Find(&groupList).Count(&countRows)
	p := paginator.NewPagintor(0, countRows)
	orm.Master().Offset(num * p.Rows).Limit(p.Rows).Order("rank_score desc,id desc").Find(&groupList)
	gconv.Structs(groupList, &groups)
	rownum := countRows / p.Rows
	return g.Map{"rownum": rownum, "total": countRows, "rows": len(groups), "data": groups}
}

func GetGroupInfo(id int) g.Map {
	//造群人数的数据
	//获取内容
	group := model.Group{}
	orm.Master().Where("id = ?", id).First(&group)
	groupMap := gconv.Map(group)
	groupQrcode := model.Qrcode{}
	rows := orm.Master().Table("qrcodes").Where("group_id = ?", id).Where("status = ?", entities.QRCODE_STATUS_NORMAL).First(&groupQrcode).RowsAffected

	if rows < 1 {
		cfg := GetSystemCfgJson()
		genGroupStatus, _ := cfg["gen_group_status"].(int)
		//活码加载优先级，最先使用数据中的活码，如果用完的
		//其次则看是否开启默认通用群.
		//最后看个人微信是否开启9,/
		if genGroupStatus == 0 {
			url, _ := cfg["default_wechat_qrcodeurl"].(string)
			groupMap["default_wechat_qrcodeurl"] = fmt.Sprintf("%s://%s%s", server.Config.Schema, server.Config.Domain, url)
		}

	} else {
		groupMap["qrcode_id"] = groupQrcode.ID
	}
	return groupMap
}

func ModifyGroupBase(id int, data g.Map) error {
	group := model.Group{}
	err := orm.Master().Table("groups").Select("id").Where("id = ?", id).First(&group).Error
	if err == gorm.ErrRecordNotFound {
		return errors.New("请确认群是否存在")
	}
	err = orm.Master().Table("groups").Where("id = ?", id).Update(data).Error
	if err != nil {
		return errors.New("系统错误")
	}
	return nil
}

func CreateGroupBase(data g.Map) (error, model.Group) {
	group := model.Group{}
	data["uniacid"] = "12"
	data["status"] = GROUP_STATUS_FREE
	title, _ := data["title"]
	rowAffected := orm.Master().Table("groups").Where("title = ?", title).First(&group).RowsAffected
	groupExits := rowAffected > 0
	if groupExits {
		return errors.New("当前名称已存在,请重新选择录入"), model.Group{}
	}
	gconv.Struct(data, &group)
	err := orm.Master().Table("groups").Create(&group).Error
	if err != nil {
		return errors.New("系统错误"), model.Group{}
	}
	return nil, group
}
