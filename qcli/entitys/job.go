package entitys

import (
	"gorm.io/gorm"
	"quanzi1/internal/front/database/migrate"
)

//迁移任务
type MigrateJob interface {
	Down(db *gorm.DB)
	Up(db *gorm.DB)
}

var Job = []MigrateJob{migrate.FrontGroupMigrate_20220930183037{}, migrate.SystemInitMigrate_20220929214331{}}
