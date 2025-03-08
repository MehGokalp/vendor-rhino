package dto

import (
	"github.com/mehgokalp/vendor-rhino/internal/entity"
	"time"
)

type Card struct {
	Balance        uint
	ActivationDate time.Time
	ExpireDate     time.Time
	Reference      string
	CardNumber     string
	Cvc            string
	Currency       entity.Currency
}
