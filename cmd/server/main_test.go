package server

import (
	"bytes"
	"encoding/json"
	"github.com/mehgokalp/vendor-rhino/internal/card/factory"
	"github.com/mehgokalp/vendor-rhino/internal/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mehgokalp/vendor-rhino/internal/delivery/http/card/create"
	"github.com/mehgokalp/vendor-rhino/internal/delivery/http/card/find"
	"github.com/mehgokalp/vendor-rhino/internal/delivery/http/card/remove"
	"github.com/mehgokalp/vendor-rhino/internal/mocks"
	"github.com/stretchr/testify/assert"
)

// TODO these are AI generated tests, they need to be reviewed and refactored

func TestFindCard(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cardRepo := &mocks.CardRepository{}
	cardRepo.On("FindByReference", "test-reference").Return(&entity.Card{Reference: "test-reference"}, nil)

	router := gin.Default()
	router.GET("/cards/:reference", find.NewHandler(cardRepo))

	t.Run("Card found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cards/test-reference", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var card entity.Card
		err := json.Unmarshal(w.Body.Bytes(), &card)
		assert.NoError(t, err)
		assert.Equal(t, "test-reference", card.Reference)
	})

	t.Run("Card not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/cards/unknown-reference", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestDeleteCard(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cardRepo := &mocks.CardRepository{}
	cardRepo.On("FindByReference", "test-reference").Return(&entity.Card{Reference: "test-reference"}, nil)

	router := gin.Default()
	router.DELETE("/cards/:reference", remove.NewHandler(cardRepo))

	t.Run("Card deleted", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/cards/test-reference", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Card not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/cards/unknown-reference", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

func TestCreateCard(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cardRepo := &mocks.CardRepository{}
	cardRepo.On("FindByReference", "test-reference").Return(&entity.Card{Reference: "test-reference"}, nil)

	currecyRepo := &mocks.CurrencyRepository{}
	currecyRepo.On("FindByCode", "USD").Return(&entity.Currency{Code: "USD"}, nil)

	router := gin.Default()
	router.POST("/cards", create.NewHandler(currecyRepo, cardRepo, factory.NewCardFactory()))

	t.Run("Card created", func(t *testing.T) {
		card := &entity.Card{Reference: "new-reference"}
		body, _ := json.Marshal(card)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/cards", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
	})

	t.Run("Card already exists", func(t *testing.T) {
		card := &entity.Card{Reference: "test-reference"}
		body, _ := json.Marshal(card)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/cards", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
	})
}
