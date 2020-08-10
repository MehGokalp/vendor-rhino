package card

import (
	"github.com/brianvoe/gofakeit/v5"
)

func populateCardInformation(card *Card) {
	gofakeit.Seed(0)
	creditCardInfo := gofakeit.CreditCard()

	card.Cvc = creditCardInfo.Cvv
	card.CardNumber = creditCardInfo.Number
	card.Reference = gofakeit.BitcoinAddress()
}
