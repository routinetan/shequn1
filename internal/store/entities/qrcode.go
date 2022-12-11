package entities

type OpQrcodeList struct {
	Id        int    `gorm:"primary_key;column:id"`
	GroupId   int    `gorm:"column:group_id"`
	QrcodeUrl string `gorm:"column:qrcode_url"`
	Title     string `gorm:"column:title"`
	Num       int    `gorm:"column:num"`
	Status    int    `gorm:"column:status"` //0为正常 1为已满
}
