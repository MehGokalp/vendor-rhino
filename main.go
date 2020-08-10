package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mehgokalp/vendor-rhino/card"
	"github.com/mehgokalp/vendor-rhino/common"
)

func main() {
	db := common.Connect()
	migrate(db)

	r := gin.Default()
	card.RegisterRoutes(r)

	// TODO: Register oauth implementation to router

	r.Run()
}

func migrate(db *gorm.DB) {
	card.AutoMigrate(db)
}
