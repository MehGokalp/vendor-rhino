package common

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var DB *gorm.DB

func Connect() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("VENDOR_RHINO_DATABASE_DNS"))

	if err != nil {
		panic(err)
	}

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

func SaveOne(obj interface{}) error {
	return DB.Save(obj).Error
}
