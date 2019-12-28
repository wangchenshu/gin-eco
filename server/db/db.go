package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Db gorm db
var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("mysql", `root:123456@tcp(127.0.0.1:3306)/eco?charset=utf8`)

	if err != nil {
		log.Panicln("err:", err.Error())
	}

	Db.SingularTable(true)
	Db.DB().SetMaxOpenConns(10)
	Db.DB().SetMaxIdleConns(10)
}
