package create

import "time"

type form struct {
	Currency       string     `form:"currency" binding:"required"`
	ActivationDate *time.Time `form:"activationDate" binding:"required" time_format:"2006-01-02"`
	ExpireDate     *time.Time `form:"expireDate" binding:"required" time_format:"2006-01-02"`
	Balance        uint       `form:"balance" binding:"required,min=0,numeric"`
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
