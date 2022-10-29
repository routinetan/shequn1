package migrate

import (
	"gorm.io/gorm"
)

type QrcodeModuleMigrate_20220930183037 struct {
	Name string
}

func (tx QrcodeModuleMigrate_20220930183037) Up(db *gorm.DB) {
	type Qrcode struct {
		ID        int    `gorm:"primary_key;column:id"`
		GroupId   int    `gorm:"column:group_id"`
		QrcodeUrl string `gorm:"column:qrcode_url"`
		Title     string `gorm:"column:title"`
		Num       int    `gorm:"column:num"`
		Status    int    `gorm:"column:status"` //0为正常 1为已满
	}
	type Group struct {
		Price    float64 `gorm:"column:price"`
		PriceTxt string  `gorm:"column:price_txt"`
		IsFree   int     `gorm:"column:is_free"`
	}
	type GroupInfo struct {
		ShowGroupNum int `gorm:"column:show_group_num"`
		CheatType    int `gorm:"column:cheat_type"`
	}
	db.AutoMigrate(&Qrcode{})
}

func (tx QrcodeModuleMigrate_20220930183037) Down(db *gorm.DB) {

}

func (tx QrcodeModuleMigrate_20220930183037) FileName() string {
	return tx.Name
}
