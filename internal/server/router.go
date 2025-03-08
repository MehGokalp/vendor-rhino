package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/internal/card/factory"
	"github.com/mehgokalp/vendor-rhino/internal/delivery/http/card/create"
	"github.com/mehgokalp/vendor-rhino/internal/delivery/http/card/find"
	"github.com/mehgokalp/vendor-rhino/internal/delivery/http/card/remove"
	"github.com/mehgokalp/vendor-rhino/internal/repository"
)

func GetRouter(
	currencyRepository *repository.CurrencyRepository,
	cardRepository *repository.CardRepository,
	cardFactory *factory.CardFactory,
) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")

	v1.POST("/card", create.NewHandler(currencyRepository, cardRepository, cardFactory))

	v1.GET("/card/:reference", find.NewHandler(cardRepository))
	v1.DELETE("/card/:reference", remove.NewHandler(cardRepository))

	return r
}
