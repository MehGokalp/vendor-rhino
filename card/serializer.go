package card

import "time"

type CardSerializer struct {
	Card
}

func (c *CardSerializer) Response() CardResponse {
	return CardResponse{
		Currency:       c.Currency.Code,
		Balance:        c.Balance,
		ActivationDate: c.ActivationDate,
		ExpireDate:     c.ExpireDate,
		Reference:      c.Reference,
		CardNumber:     c.CardNumber,
		Cvc:            c.Cvc,
	}
}

type CardResponse struct {
	Currency       string     `json:"currency"`
	Balance        uint       `json:"balance"`
	ActivationDate *time.Time `json:"activation_date"`
	ExpireDate     *time.Time `json:"expire_date"`
	Reference      string     `json:"reference"`
	CardNumber     string     `json:"card_number"`
	Cvc            string     `json:"cvc"`
}
