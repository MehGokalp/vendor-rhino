package card

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/common"
	"net/http"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/create", createCard)
	r.GET("/find/:reference", findCard)
	r.DELETE("/remove", removeCard)
}

func createCard(c *gin.Context) {
	validator := NewCardValidator()

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
	// TODO: Implement find card action
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func removeCard(c *gin.Context) {
	// TODO: Implement remove card action
	c.JSON(http.StatusNoContent, gin.H{})
}
