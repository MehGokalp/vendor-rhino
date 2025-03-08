package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/mehgokalp/vendor-rhino/internal/entity"
)

type CardRepositoryInterface interface {
	Create(card *entity.Card) error
	FindByReference(reference string) *entity.Card
	Delete(card *entity.Card) error
}

type CardRepository struct {
	db *gorm.DB
}

func NewCardRepository(db *gorm.DB) *CardRepository {
	return &CardRepository{db}
}

func (r *CardRepository) Create(card *entity.Card) error {
	return r.db.Create(card).Error
}

func (r *CardRepository) FindByReference(reference string) *entity.Card {
	card := entity.Card{}
	r.db.Where("reference = ?", reference).Preload("Currency").First(&card)

	return &card
}

func (r *CardRepository) Delete(card *entity.Card) error {
	return r.db.Delete(card).Error
}
