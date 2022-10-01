package main

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"quanzi1/foundation/util"
	migrate2 "quanzi1/qcli/entitys"
	"quanzi1/qcli/lib/migrate"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/quanzi1?charset=utf8mb4&parseTime=True&loc=Local"
	dbhandle, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	migrationPath := "/data/migration"
	if util.Exists(migrationPath) == false {
		os.MkdirAll(migrationPath, 0777)
	}

	infop, err := os.OpenFile(migrationPath+"/db_info.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	metap, err := os.OpenFile(migrationPath+"/db.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		panic(err)
	}

	var head []migrate.MigrateHead
	bhead, _ := ioutil.ReadAll(infop)
	json.Unmarshal(bhead, &head)

	var infolist []migrate.MigrateList
	binfo, _ := ioutil.ReadAll(metap)
	json.Unmarshal(binfo, &infolist)

	for _, v := range migrate2.Job {
		v.Up(dbhandle)
	}
}
