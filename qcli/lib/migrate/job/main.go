package main

import (
	"encoding/json"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"os"
	"shequn1/foundation/util"
	migrate2 "shequn1/qcli/entitys"
	"shequn1/qcli/lib/migrate"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	dsn := "root:root@tcp(127.0.0.1:3306)/quanzi1?charset=utf8mb4&parseTime=True&loc=Local"
	dbhandle, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	migrationPath := "data/migration"
	if util.Exists(migrationPath) == false {
		os.MkdirAll(migrationPath, 0777)
	}

	infop, err := os.OpenFile(migrationPath+"/db_info.json", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	metap, err := os.OpenFile(migrationPath+"/db.json", os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}

	var head migrate.MigrateHead
	bhead, err := ioutil.ReadAll(metap)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(bhead, &head)

	var infolist []migrate.MigrateList
	binfo, err := ioutil.ReadAll(infop)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(binfo, &infolist)
	head.CurVersion += 1
	for _, v := range migrate2.Job {
		temp := migrate.MigrateList{}
		temp.Batch = head.CurVersion
		temp.Name = v.FileName()
		temp.UpdataAt = time.Now().Format("2006-01-02T15:04:05")
		v.Up(dbhandle)
		infolist = append(infolist, temp)
	}

	head.LastTime = time.Now().Format("2006-01-02T15:04:05")
	head.ID = 1
	shead, err := json.Marshal(head)
	if err != nil {
		panic(err)
	}
	metap.Seek(0, io.SeekStart)
	_, err = metap.WriteString(string(shead))
	if err != nil {
		panic(err)
	}
	sinfolist, err := json.Marshal(infolist)
	if err != nil {
		panic(err)
	}
	infop.Seek(0, io.SeekStart)
	_, err = infop.WriteString(string(sinfolist))
	if err != nil {
		panic(err)
	}
}
