package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zhaoqingsong/studyweb/model"
	"log"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/filetest.db")
	log.Println("打开数据库成功")
	if err != nil {
		log.Println(fmt.Sprintf("【initbaseDB.NewEngine】ex:%s\n", err.Error()))
		os.Exit(0)
		return
	}
	//defer db.Close()
	// Migrate the schema
	db.AutoMigrate(&model.File{})
	err = db.DB().Ping()
	if err != nil {
		log.Println(fmt.Sprintf("【initbaseDB.Ping】ex:%s\n", err.Error()))
		os.Exit(0)
		return
	}
	//db.DB().SetMaxIdleConns(config.Config.DB.MaxIdleConn)
	//db.DB().SetMaxOpenConns(config.Config.DB.MaxOpenConn)

}
