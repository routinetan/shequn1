package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"quanzi1/foundation/database/orm"
	"quanzi1/foundation/paginator"
	"quanzi1/internal/entities"
	"quanzi1/internal/model"
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
