package migrate

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"shequn1/foundation/util"
)

type MigrateJob struct {
}

type MigrateHead struct {
	ID         int
	CurVersion int
	LastTime   string
}

type MigrateList struct {
	ID       int
	Name     string
	Batch    int
	Author   string
	UpdataAt string
}

func GetMigrateInfo() (MigrateHead, []MigrateList) {
	migrationPath := "data/migration"
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
	var head MigrateHead
	bhead, _ := ioutil.ReadAll(metap)
	json.Unmarshal(bhead, &head)

	var infolist []MigrateList
	binfo, _ := ioutil.ReadAll(infop)
	json.Unmarshal(binfo, &infolist)
	return head, infolist
}
