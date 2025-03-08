package find

import "time"

type form struct {
	Reference string `uri:"reference" binding:"required"`
}

type response struct {
	Currency       string     `json:"currency"`
	Balance        uint       `json:"balance"`
	ActivationDate *time.Time `json:"activation_date"`
	ExpireDate     *time.Time `json:"expire_date"`
	Reference      string     `json:"reference"`
	CardNumber     string     `json:"card_number"`
	Cvc            string     `json:"cvc"`
}
