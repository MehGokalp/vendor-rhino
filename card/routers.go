package card

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/create", createCard)
	r.GET("/find", findCard)
	r.DELETE("/remove", removeCard)
}

func createCard(c *gin.Context) {
	// TODO: Implement create card action
	c.JSON(http.StatusCreated, gin.H{
		"message": "OK",
	})
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
