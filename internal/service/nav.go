package service

import (
	"github.com/gogf/gf/util/gconv"
	"quanzi1/foundation/database/orm"
	"quanzi1/internal/model"
)

func GetNavList() []string {
	var nav []model.Category
	orm.Master().Find(&nav).Where(&model.Category{Status: 1})
	gconv.Structs()
	return []string{}
}
