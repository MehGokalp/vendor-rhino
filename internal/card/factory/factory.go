package factory

import (
	"github.com/mehgokalp/vendor-rhino/internal/card/dto"
	"github.com/mehgokalp/vendor-rhino/internal/entity"
)

type CardFactory struct {
}

func NewCardFactory() *CardFactory {
	return &CardFactory{}
}

func (f *CardFactory) Create(card dto.Card) entity.Card {
	return entity.Card{
		CardNumber:     card.CardNumber,
		Cvc:            card.Cvc,
		Reference:      card.Reference,
		Balance:        card.Balance,
		ActivationDate: &card.ActivationDate,
		ExpireDate:     &card.ExpireDate,
		Currency:       &card.Currency,
	}
}
