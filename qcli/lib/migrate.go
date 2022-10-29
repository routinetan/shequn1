package lib

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {

	dsn := "root:root@tcp(127.0.0.1:3306)/shequn1?charset=utf8mb4&parseTime=True&loc=Local"
	dbhandle, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m := gormigrate.New(dbhandle, gormigrate.DefaultOptions, []*gormigrate.Migration{{
		ID: "201608301400",
		Migrate: func(tx *gorm.DB) error {
			// it's a good pratice to copy the struct inside the function,
			// so side effects are prevented if the original struct changes during the time
			type Person struct {
				gorm.Model
				Name string
			}
			return tx.AutoMigrate(&Person{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable("people")
		},
	}})
	if err = m.Migrate(); err != nil {
		log.Fatalf("Could not migrateJob: %v", err)
	}
	log.Printf("Migration did run successfully")
}
