package card

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/common"
	"time"
)

type CardValidator struct {
	Input struct {
		Currency       string     `form:"currency" binding:"required"`
		ActivationDate *time.Time `form:"activationDate" binding:"required,datetime=2020-05-10"`
		ExpireDate     *time.Time `form:"expireDate" binding:"required,datetime=2020-05-10"`
		Balance        uint       `form:"balance" binding:"required,min=0,numeric"`
	}
	Model Card
}

func (v *CardValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}

	v.Model.Currency = Currency{
		Code: v.Input.Currency,
	}
	v.Model.ActivationDate = v.Input.ActivationDate
	v.Model.ExpireDate = v.Input.ExpireDate
	v.Model.Balance = v.Input.Balance

	return nil
}

func NewCardValidator() CardValidator {
	return CardValidator{}
}
