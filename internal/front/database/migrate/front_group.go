package migrate

import (
	"fmt"
	"gorm.io/gorm"
)

var cateNames = []string{
	"美业群", "行业群", "高校群", "求职招聘群", "相亲交友群", "微商群", "联盟群", "二手车群",
	"闲货处理群", "宝妈群", "创业群", "其他群", "其他互助群", "红包群", "兼职群", "拼多多互助群", "资源互换群",
}

type FrontGroupMigrate_20220930183037 struct {
}

func (tx FrontGroupMigrate_20220930183037) Up(db *gorm.DB) {
	for _, catename := range cateNames {
		sql := fmt.Sprintf("INSERT INTO categories (`name`) VALUES ('%s');", catename)
		db.Raw(sql)
	}
	fmt.Printf("SystemInitMigrate up")
}

func (tx FrontGroupMigrate_20220930183037) Down(db *gorm.DB) {

}
