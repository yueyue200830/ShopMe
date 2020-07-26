// Database package.
// Connect to database and search for results
package dao

import (
	"backend-go/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	dbType, conn := conf.GetDatabaseSetting()
	db, err = gorm.Open(dbType, conn)
	if err != nil {
		fmt.Println(err)
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
