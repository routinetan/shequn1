package model

type GroupInfo struct {
	GroupId      int                    `gorm:"column:group_id"`
	Content      string                 `gorm:"column:content"`
	LabelContent string                 `gorm:"content:label_content"`
	Number       string                 `gorm:"column:number"`
	Extra        map[string]interface{} `gorm:"column:extra"`
	DisplayMode  int                    `gorm:"display_mode"`
	QrCode       string                 `gorm:"qr_code"`
}
