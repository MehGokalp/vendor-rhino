package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:123456@/vendor_rhino?charset=utf8&parseTime=True")

	if err != nil {
		panic(err)
	}

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}