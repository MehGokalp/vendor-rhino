package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/mehgokalp/vendor-rhino/internal/entity"
)

func LoadFixtures(db *gorm.DB) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, code := range []string{"USD", "EUR", "GBP"} {
		if err := tx.Create(&entity.Currency{Code: code}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}
