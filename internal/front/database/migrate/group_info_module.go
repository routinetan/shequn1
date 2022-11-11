package migrate

import (
	"gorm.io/gorm"
)

type GroupInfoModule_202211111708 struct {
	Name string
}

func (tx GroupInfoModule_202211111708) Up(db *gorm.DB) {
	db.AutoMigrate(&Qrcode{})
}

func (tx GroupInfoModule_202211111708) Down(db *gorm.DB) {

}

func (tx GroupInfoModule_202211111708) FileName() string {
	return tx.Name
}
