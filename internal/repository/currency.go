package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/mehgokalp/vendor-rhino/internal/entity"
)

type CurrencyRepositoryInterface interface {
	FindByCode(code string) *entity.Currency
}

type CurrencyRepository struct {
	db *gorm.DB
}

func NewCurrencyRepository(db *gorm.DB) *CurrencyRepository {
	return &CurrencyRepository{db}
}

func (r *CurrencyRepository) FindByCode(code string) *entity.Currency {
	currency := entity.Currency{}
	r.db.First(&currency, "code = ?", code)

	return &currency
}
