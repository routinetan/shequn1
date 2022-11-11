package migrate

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SystemCfgModuleMigrate_2022110832123 struct {
	Name string
}

func (tx SystemCfgModuleMigrate_2022110832123) Up(db *gorm.DB) {
	type SystemCfg struct {
		Id        int            `gorm:"primary_key;column:id"`
		Label     string         `gorm:"label"`
		Cfg       datatypes.JSON `gorm:"cfg"`
		CreatedAt string         `gorm:"created_at"`
		UpdateAt  string         `gorm:"updated_at"`
	}
	db.AutoMigrate(&SystemCfg{})
}

func (tx SystemCfgModuleMigrate_2022110832123) Down(db *gorm.DB) {

}

func (tx SystemCfgModuleMigrate_2022110832123) FileName() string {
	return tx.Name
}
