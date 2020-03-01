package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// Db gorm db
var Db *gorm.DB

func init() {
	var err error
	godotenv.Load("../../.env")
	myEnv, err := godotenv.Read()
	connString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8",
		myEnv["DATABASE_USER"],
		myEnv["DATABASE_PASSWORD"],
		myEnv["DATABASE_HOST"],
		myEnv["DATABASE_PORT"],
		myEnv["DATABASE_DB"],
	)
	Db, err = gorm.Open("mysql", connString)
	if err != nil {
		log.Panicln("err:", err.Error())
	}

	Db.SingularTable(true)
	Db.DB().SetMaxOpenConns(10)
	Db.DB().SetMaxIdleConns(10)
}
