package service

import (
	"github.com/gogf/gf/util/gconv"
	"shequn1/foundation/database/orm"
	"shequn1/internal/entities"
	"shequn1/internal/model"
)

func GetNavList() []entities.Category {
	var navEntity []entities.Category
	var nav []model.Category
	orm.Master().Where(&model.Category{Status: 1}).Find(&nav)
	gconv.Structs(nav, &navEntity)
	return navEntity
}
