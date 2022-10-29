package migrate

import (
	"gorm.io/gorm"
	"time"
)

type AsystemInitMigrate_20220929214331 struct {
	Name string
}

func (tx AsystemInitMigrate_20220929214331) Up(db *gorm.DB) {
	type Group struct {
		ID        int    `gorm:"primary_key;column:id"`
		Title     string `gorm:"column:title"`
		UniacId   int    `gorm:"column:uniacid"`
		ClassId   int    `gorm:"column:classid"`
		QrCode    int    `gorm:"column:qrcode"`
		Label     string `gorm:"column:label"`
		Type      int    `gorm:"column:type"`
		RankScore int    `gorm:"column:rank_score"`
		IsShow    int    `gorm:"column:is_show"`
		Status    int    `gorm:"colum:status"`
		Tag       string `gorm:"colum:tag"`
		Thumb     string `gorm:"thumb"`
	}
	type GroupInfo struct {
		GroupId      string `gorm:"column:group_id"`
		Content      string `gorm:"column:content"`
		LabelContent string `gorm:"content:label_content"`
		Number       string `gorm:"column:number"`
		Extra        string `gorm:"column:extra"`
		DisplayMode  int    `gorm:"display_mode"`
		QrCode       string `gorm:"qr_code"`
	}
	type QuickLink struct {
		ID     int    `gorm:primary_key,column:"id"`
		Icon   string `gorm:"column:icon"`
		Label  string `gorm:"column:label"`
		Path   string `gorm:"column:path"`
		Status string `gorm:"column:status"`
	}
	type Category struct {
		ID          int    `gorm:"primary_key;column:id;" bson:"_id"`
		Uniacid     int    `gorm:"column:uniacid"`
		CreatedAt   int    `gorm:"column:created_at;index:created_at" bson:"created_at"`
		UpdatedAt   int    `gorm:"column:updated_at;index:updated_at" bson:"updated_at"`
		Alias       string `gorm:"column:alias" bson:"alias"`
		Name        string `gorm:"column:name" bson:"name"`
		Pic         string `gorm:"column:pic" bson:"pic"`
		Badge       string `gorm:"column:badge" bson:"badge"`
		Description string `gorm:"column:description" bson:"description"`
		Pid         string `gorm:"column:pid" bson:"pid"`
		Status      string `gorm:"status"`
		Level       int    `gorm:"column:level" bson:"level"`
	}
	type UserDemand struct {
	}
	type UserDemandRadar struct {
		Id               int
		UserId           int
		OneParamsId      int
		OneParamsValue   string
		TwoParamsId      int
		TwoParamsValue   string
		ThreeParamsId    int
		ThreeParamsValue string
		FourParamId      int
		FourParamsValue  string
		FiveParamId      int
		FiveParamsValue  string
	}
	type UserSocialRelation struct {
		Id        int
		OneUserId int
		TwoUserId int
		Type      int
		Level     int
		Status    int
	}
	type UserCompagin struct {
		Id           int
		UserId       int
		CompaginId   int
		ActiveCost   float64
		DownloadCost float64
	}
	type UserTube struct { //用户连接表,用于帮助用户和系统建立关系
		Id        int
		UserId    int
		Type      int
		CreatedAt time.Time
	}
	type TubeUserHintActivity struct {
		Id      int
		UserId  int
		NotesId int
		HintsAt time.Time
		Type    int
	}
	db.AutoMigrate(&Group{}, &GroupInfo{}, &Category{}, &QuickLink{})

}

func (tx AsystemInitMigrate_20220929214331) Down(db *gorm.DB) {
	db.Migrator().DropTable("")
}

func (tx AsystemInitMigrate_20220929214331) FileName() string {
	return tx.Name
}
