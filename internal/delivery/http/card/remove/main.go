package remove

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/internal/repository"
	pkgErrors "github.com/mehgokalp/vendor-rhino/pkg/errors"
	"net/http"
)

type Handler struct {
	repository repository.CardRepositoryInterface
}

func NewHandler(repository repository.CardRepositoryInterface) func(*gin.Context) {
	return func(c *gin.Context) {
		h := Handler{
			repository: repository,
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

	card := h.repository.FindByReference(submittedForm.Reference)
	if card == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, errors.New("card not found"))
		return
	}

	err = h.repository.Delete(card)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, errors.New("error on remove"))
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
