package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"shequn1/foundation/database/orm"
	"shequn1/foundation/paginator"
	"shequn1/internal/entities"
	"shequn1/internal/model"
)

func GetGroupList(num int) g.Map {
	var groups []entities.Group
	var groupList []model.Group
	var countRows int
	orm.Master().Find(&groupList).Count(&countRows)
	p := paginator.NewPagintor(0, countRows)
	orm.Master().Offset(num).Limit(p.Rows).Order("rank_score desc").Find(&groupList)
	gconv.Structs(groupList, &groups)
	rownum := countRows / p.Rows
	return g.Map{"rownum": rownum + 1, "rows": len(groups), "data": groups}
}

func GetOpGroupList(num int) g.Map {
	var groups []entities.Group
	var groupList []model.Group
	var countRows int
	orm.Master().Find(&groupList).Count(&countRows)
	p := paginator.NewPagintor(0, countRows)
	orm.Master().Offset(num * p.Rows).Limit(p.Rows).Order("rank_score desc").Find(&groupList)
	gconv.Structs(groupList, &groups)
	rownum := countRows / p.Rows
	return g.Map{"rownum": rownum + 1, "rows": len(groups), "data": groups}
}

func GetGroupInfo(id int) g.Map {
	//造群人数的数据
	//获取内容
	return g.Map{}
}
