package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var GormDb *gorm.DB

func Init() {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASS")
	PORT := os.Getenv("DB_PORT")
	HOST := os.Getenv("DB_HOST")
	DB := os.Getenv("DB_NAME")
	args := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USER, PASS, HOST, PORT, DB)
	db, err := gorm.Open("mysql", args)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database with setting: %s", args))
	}
	GormDb = db
	fmt.Print("Connected to mysql")
	pingTicker := time.NewTicker(30 * time.Second)
	pingDone := make(chan bool)
	go func() {
		for {
			select {
			case <-pingDone:
				return
			case <-pingTicker.C:
				b := pingDb(GormDb.DB())
				if !b {
					pingDone <- true
				}
			}
		}
	}()

}
func pingDb(db *sql.DB) bool {
	er := db.Ping()
	if er != nil {
		return false
	} else {
		return true
	}
}
