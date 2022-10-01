package main

import (
	"fmt"
	"github.com/gogf/gf/container/glist"
	"github.com/gogf/gf/container/gset"
	"io/ioutil"
	"os"
	"quanzi1/qcli/lib/migrate"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func SkipFile(src, dst *gset.Set) *glist.List {
	return glist.NewFrom(src.Diff(dst).Slice())
}
func SaveMigrateResult() []string {
	return []string{}
}

// 把秒级的时间戳转为time格式
func SecondToTime(sec int64) time.Time {
	return time.Unix(sec, 0)
}

func GetCreateTime(file os.FileInfo) int64 {
	if runtime.GOOS == "windows" {
		winFileAttr := file.Sys().(*syscall.Win32FileAttributeData)
		return SecondToTime(winFileAttr.CreationTime.Nanoseconds() / 1e9).Unix()
		//linuxFileAttr := file.Sys().(*syscall.Stat_t)
	} else if runtime.GOOS == "linux" {
		//linuxFileAttr := file.Sys().(*syscall.Stat_t)
	}
	return 0
}
func main() {
	head, infolist := migrate.GetMigrateInfo()
	mode := "front"
	migrationPath := fmt.Sprintf("e:/goproject/quanzi1/internal/%s/database/migrate", mode)
	files, _ := ioutil.ReadDir(migrationPath)
	migrateFile := gset.New(true)
	skipedFile := gset.New(true)
	for _, v := range files {
		migrateFile.Add(v.Name())
	}
	for _, v := range infolist {
		skipedFile.Add(v)
	}
	fmt.Println(head.CurVersion)
	//跳过已经处理过的问价
	processfile := SkipFile(migrateFile, skipedFile)
	var migrations []string
	processfile.Iterator(func(e *glist.Element) bool {
		filepath, _ := e.Value.(string)
		filename := filepath
		pos := strings.Index(filename, ".go")
		filename = filepath[:pos]
		filenames := strings.Split(filename, "_")
		for k, v := range filenames {
			filenames[k] = strings.ToUpper(string(v[0])) + v[1:]
		}
		jobName := strings.Join(filenames, "")
		info, _ := os.Stat(migrationPath + "/" + filepath)
		unixAt := GetCreateTime(info)
		createAt := time.Unix(unixAt, 0).Format("20060102150405")
		//添加已经处理过的文件记录
		migrateName := fmt.Sprintf("migrate.%sMigrate_%s{}", jobName, createAt)
		migrations = append(migrations, migrateName)
		return true
	})
	jobs := strings.Join(migrations, ",")
	jobcode := fmt.Sprintf(`package entitys
import (
	"gorm.io/gorm"
	"quanzi1/internal/front/database/migrate"
)
//迁移任务
type MigrateJob interface {
	Down(db *gorm.DB)
	Up(db *gorm.DB)
}

var Job = []MigrateJob{%s}`, jobs)

	fp, err := os.OpenFile("../entitys/job.go", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("open error", err)
		return
	}
	_, err = fp.WriteString(jobcode)
	if err != nil {
		fmt.Println("write error", err)
		return
	}
}
