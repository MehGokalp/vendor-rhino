package card

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/mehgokalp/vendor-rhino/common"
)

func populateCardInformation(card *Card) {
	creditCardInfo := gofakeit.CreditCard()
	card.Cvc = creditCardInfo.Cvv
	card.CardNumber = creditCardInfo.Number
	card.Reference = common.RandStringBytes(12)
}
