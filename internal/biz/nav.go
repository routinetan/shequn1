package biz

import (
	"github.com/gogf/gf/util/gconv"
	"shequn1/internal/foundation/database/orm"
	"shequn1/internal/store/entities"
	"shequn1/internal/store/model"
)

func GetNavList() []entities.Category {
	var navEntity []entities.Category
	var nav []model.Category
	orm.Master().Where(&model.Category{Status: 1}).Find(&nav)
	gconv.Structs(nav, &navEntity)
	return navEntity
}
