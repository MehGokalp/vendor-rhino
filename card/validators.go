package card

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/common"
	"time"
)

type CardValidator struct {
	Input struct {
		Currency       string     `form:"currency" binding:"required"`
		ActivationDate *time.Time `form:"activationDate" binding:"required" time_format:"2006-01-02"`
		ExpireDate     *time.Time `form:"expireDate" binding:"required" time_format:"2006-01-02"`
		Balance        uint       `form:"balance" binding:"required,min=0,numeric"`
	}
	Model Card
}

func (v *CardValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}

	var currency Currency
	DB := common.GetDB()
	DB.Where("code = ?", v.Input.Currency).First(&currency)

	if currency.ID == 0 {
		return CurrencyNotFoundError{currencyCode: v.Input.Currency}
	}

	v.Model.Currency = &currency
	v.Model.ActivationDate = v.Input.ActivationDate
	v.Model.ExpireDate = v.Input.ExpireDate
	v.Model.Balance = v.Input.Balance

	return nil
}

func NewCardValidator() CardValidator {
	return CardValidator{}
}
