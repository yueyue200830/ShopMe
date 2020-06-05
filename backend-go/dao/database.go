package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@(localhost)/go_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connect to database successfully")
}

func Close() {
	err := db.Close()
	if err != nil {
		panic("failed to close database")
	}
}

func GetDB() *gorm.DB {
	return db
}
