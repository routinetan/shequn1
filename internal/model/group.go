package model

type Group struct {
	ID        int      `gorm:"primary_key;column:id"`
	Title     string   `gorm:"column:title"`
	UniacId   int      `gorm:"column:uniacid"`
	ClassId   int      `gorm:"column:classid"`
	QrCode    int      `gorm:"column:qrcode"`
	Label     string   `gorm:"column:label"`
	Type      int      `gorm:"column:type"`
	RankScore int      `gorm:"column:rank_score"`
	IsShow    int      `gorm:"column:is_show"`
	Status    int      `gorm:"colum:status"`
	Tag       []string `gorm:"colum:tag"`
	Thumb     string   `gorm:"thumb"`
}
