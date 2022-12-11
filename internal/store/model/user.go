package model

type User struct {
	Id         string `gorm:"column:id"`
	Indentifed string `gorm:"column:indentifed"`
	Mobile     string `gorm:"column:mobile"`
	Account    string `gorm:"column:account"`
	Pwd        string `gorm:"columns:pwd"`
}
