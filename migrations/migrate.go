package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/mehgokalp/vendor-rhino/internal/entity"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&entity.Card{})
	db.AutoMigrate(&entity.Currency{})
}
