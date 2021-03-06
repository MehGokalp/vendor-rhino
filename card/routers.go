package card

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/common"
	"net/http"
)

func RegisterRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")

	v1.POST("/", createCard)
	v1.GET("/:reference", findCard)
	v1.DELETE("/:reference", deleteCard)
}

func createCard(c *gin.Context) {
	validator := NewCreateCardValidator()

	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	populateCardInformation(&validator.Model)

	if err := common.SaveOne(&validator.Model); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewError("database", err))
		return
	}

	serializer := CardSerializer{validator.Model}
	c.JSON(http.StatusCreated, serializer.Response())
}

func findCard(c *gin.Context) {
	validator := NewFindCardValidator()

	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	serializer := CardSerializer{validator.Model}
	c.JSON(http.StatusOK, serializer.Response())
}

func deleteCard(c *gin.Context) {
	validator := NewDeleteCardValidator()

	if err := validator.Bind(c); err != nil {
		c.JSON(http.StatusUnprocessableEntity, common.NewValidatorError(err))
		return
	}

	common.GetDB().Delete(validator.Model)

	c.JSON(http.StatusNoContent, gin.H{})
}
