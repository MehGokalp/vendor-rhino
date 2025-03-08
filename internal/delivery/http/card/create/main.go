package create

import (
	"errors"
	"github.com/brianvoe/gofakeit/v5"
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/internal/card/dto"
	"github.com/mehgokalp/vendor-rhino/internal/card/factory"
	"github.com/mehgokalp/vendor-rhino/internal/repository"
	pkgErrors "github.com/mehgokalp/vendor-rhino/pkg/errors"
	"net/http"
)

type Handler struct {
	currencyRepository repository.CurrencyRepositoryInterface
	cardRepository     repository.CardRepositoryInterface
	factory            *factory.CardFactory
}

func NewHandler(currencyRepository repository.CurrencyRepositoryInterface, cardRepository repository.CardRepositoryInterface, factory *factory.CardFactory) func(*gin.Context) {
	return func(c *gin.Context) {
		h := Handler{
			currencyRepository: currencyRepository,
			cardRepository:     cardRepository,
			factory:            factory,
		}

		h.Handle(c)
	}
}

func (h *Handler) Handle(c *gin.Context) {
	var submittedForm form
	err := c.Bind(&submittedForm)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, pkgErrors.NewHttpValidationError(err))
		return
	}

	currency := h.currencyRepository.FindByCode(submittedForm.Currency)
	if currency.ID == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.New("currency not found"))
		return
	}

	gofakeit.Seed(0)
	creditCardInfo := gofakeit.CreditCard()

	card := h.factory.Create(dto.Card{
		Balance:        submittedForm.Balance,
		ActivationDate: *submittedForm.ActivationDate,
		ExpireDate:     *submittedForm.ExpireDate,
		Reference:      gofakeit.BitcoinAddress(),
		CardNumber:     creditCardInfo.Number,
		Cvc:            creditCardInfo.Cvv,
		Currency:       *currency,
	})

	if err := h.cardRepository.Create(&card); err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}

	c.JSON(http.StatusCreated, response{
		Currency:       currency.Code,
		Balance:        card.Balance,
		ActivationDate: card.ActivationDate,
		ExpireDate:     card.ExpireDate,
		Reference:      card.Reference,
		CardNumber:     card.CardNumber,
		Cvc:            card.Cvc,
	})
}
