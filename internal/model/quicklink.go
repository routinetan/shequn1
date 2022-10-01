package model

type QuickLink struct {
	ID    int    `gorm:primary_key,column:"id"`
	Icon  string `gorm:"column:icon"`
	Label string `gorm:"column:label"`
}
