package entities

type QrcodeForm struct {
	Title     string `form:"title"     json:"title"`
	Status    int    `form:"status"    json:"status"`
	QrcodeUrl string `form:"thumb_url" json:"qrcode_url"`
	GroupId   int    `form:"group_id"  json:"group_id"`
}
