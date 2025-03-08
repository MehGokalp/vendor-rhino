package migrations

import "github.com/jinzhu/gorm"

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
