package model

type SystemSetting struct {
	Title                  string `json:"title"`
	DefaultWechatStatus    int    `json:"default_wechat_status"`
	DefaultWechatQrCodeUrl string `json:"default_wechat_qrcodeurl"`
	VipStatus              int    `json:"vip_status"`
	GenGroupStatus         int    `json:"gen_group_status"`
	PayCloseStatus         int    `json:"pay_close_status"`
	FeederDomainStatus     int    `json:"feeder_domain_status"`
	GenGroupId             int    `json:"gen_group_id"`
	CloseStatus            int    `json:"close_status"`
}

type SystemCfg struct {
	Id        int
	Label     string
	Cfg       string
	CreatedAt string
	UpdateAt  string
}
