package card

import (
	"github.com/jinzhu/gorm"
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
		if err := tx.Create(&Currency{Code: code}).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func FlushDB(db *gorm.DB) error {
	var tableInfo []struct {
		Table string
	}
	query := `SELECT table_name "table"
				FROM information_schema.tables WHERE table_schema='vendor_rhino'
					AND table_type='BASE TABLE';`
	db.Raw(query).Scan(&tableInfo)

	dropQueries := make([]string, len(tableInfo))

	for i, info := range tableInfo {
		dropQueries[i] = "DROP TABLE " + info.Table + ";"
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, query := range dropQueries {
		tx.Exec(query)
	}

	return tx.Commit().Error
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&Currency{})
}
