package migrate

import (
	"gorm.io/gorm"
)

type {{.Name}} struct {
	Name string
}

func (tx {{.Name}}) Up(db *gorm.DB) {
	db.AutoMigrate(&Qrcode{})
}

func (tx {{.Name}}) Down(db *gorm.DB) {

}

func (tx {{.Name}}) FileName() string {
	return tx.Name
}
