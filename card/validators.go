package card

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/common"
	"time"
)

// BEGIN - Create card action validator

type CreateCardValidator struct {
	Input struct {
		Currency       string     `form:"currency" binding:"required"`
		ActivationDate *time.Time `form:"activationDate" binding:"required" time_format:"2006-01-02"`
		ExpireDate     *time.Time `form:"expireDate" binding:"required" time_format:"2006-01-02"`
		Balance        uint       `form:"balance" binding:"required,min=0,numeric"`
	}
	Model Card
}

func (v *CreateCardValidator) Bind(c *gin.Context) error {
	err := common.Bind(c, v)
	if err != nil {
		return err
	}

	currency := Currency{}
	DB := common.GetDB()
	DB.First(&currency, "code = ?", v.Input.Currency)

	if currency.ID == 0 {
		return CurrencyNotFoundError{currencyCode: v.Input.Currency}
	}

	v.Model.Currency = &currency
	v.Model.ActivationDate = v.Input.ActivationDate
	v.Model.ExpireDate = v.Input.ExpireDate
	v.Model.Balance = v.Input.Balance

	return nil
}

func NewCreateCardValidator() CreateCardValidator {
	return CreateCardValidator{}
}

// END - Create card action validator

// BEGIN - Find card action validator

type FindCardValidator struct {
	Input struct {
		Reference string `uri:"reference" binding:"required"`
	}
	Model Card
}

func (v *FindCardValidator) Bind(c *gin.Context) error {
	err := common.BindUrl(c, v)
	if err != nil {
		return err
	}

	var card Card
	DB := common.GetDB()
	DB.Where("reference = ?", v.Input.Reference).Preload("Currency").First(&card)

	if card.ID == 0 {
		return CardNotFoundError{reference: v.Input.Reference}
	}

	v.Model = card

	return nil
}

func NewFindCardValidator() FindCardValidator {
	return FindCardValidator{}
}

// END - Find card action validator

type DeleteCardValidator struct {
	Input struct {
		Reference string `uri:"reference" binding:"required"`
	}
	Model Card
}

func (v *DeleteCardValidator) Bind(c *gin.Context) error {
	err := common.BindUrl(c, v)
	if err != nil {
		return err
	}

	var card Card
	DB := common.GetDB()
	DB.Where("reference = ?", v.Input.Reference).First(&card)

	if card.ID == 0 {
		return CardNotFoundError{reference: v.Input.Reference}
	}

	v.Model = card

	return nil
}

func NewDeleteCardValidator() DeleteCardValidator {
	return DeleteCardValidator{}
}
