package card

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Card struct {
	gorm.Model
	ID             uint       `gorm:"column:id"`
	Balance        uint       `gorm:"column:balance"`
	ActivationDate *time.Time `gorm:"column:activation_date"`
	ExpireDate     *time.Time `gorm:"column:expire_date"`
	Reference      string     `gorm:"size:255;unique_index;column:reference"`
	CardNumber     string     `gorm:"size:255;unique_index;column:card_number"`
	Cvc            string     `gorm:"size:255;column:cvc"`
	Active         bool       `gorm:"column:active,default:true"`
	CurrencyId     uint       `gorm:"column:currency_id"`
	Currency       *Currency
}

func (c *Card) TableName() string {
	return "card"
}

type Currency struct {
	ID    uint
	Code  string `gorm:"unique_index"`
	Cards []Card `gorm:"ForeignKey:CurrencyId"`
}

func (c *Currency) TableName() string {
	return "currency"
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&Currency{})
}
