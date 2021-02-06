package card

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mehgokalp/vendor-rhino/common"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

var router *gin.Engine

func getRouter() *gin.Engine {
	if router != nil {
		return router
	}

	router = gin.Default()
	RegisterRoutes(router)

	return router
}

func processRequest(r *gin.Engine, req *http.Request, t *testing.T, statusCode int) []byte {
	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	data, err := ioutil.ReadAll(w.Body)
	assert.Nil(t, err)

	if assert.Equal(t, statusCode, w.Code) != true {
		err := fmt.Sprintf("An error occurred: '%s'", string(data))
		t.Errorf(err)
		t.Fail()
	}

	return data
}

func openDBConnection(t *testing.T) *gorm.DB {
	db := common.Connect()
	if err := FlushDB(db); err != nil {
		t.Errorf("there were error while flushing the db: %s", err)
	}

	AutoMigrate(db)

	if err := LoadFixtures(db); err != nil {
		t.Errorf("there were error while loading fixtures: %s", err)
	}

	return db
}

func createTestCard(t *testing.T) CardResponse {
	r := getRouter()
	activationDate := time.Now().Add(time.Hour * 100)
	expireDate := time.Now().Add(time.Hour * 150)
	currency := "EUR"

	payload := url.Values{}
	payload.Set("currency", currency)
	payload.Set("activationDate", activationDate.Format("2006-01-02"))
	payload.Set("expireDate", expireDate.Format("2006-01-02"))
	payload.Set("balance", "5000")

	req, err := http.NewRequest(http.MethodPost, "/v1/", strings.NewReader(payload.Encode()))
	assert.Nil(t, err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	data := processRequest(r, req, t, http.StatusCreated)

	response := CardResponse{}

	err = json.Unmarshal(data, &response)
	assert.Nil(t, err)

	assert.Equal(t, currency, response.Currency)
	assert.Equal(t, uint(5000), response.Balance)

	assert.NotNil(t, response.ActivationDate)
	assert.NotNil(t, response.ExpireDate)

	assert.NotNil(t, response.Reference)
	assert.NotNil(t, response.Cvc)
	assert.NotNil(t, response.CardNumber)

	return response
}

func TestCreateCard(t *testing.T) {
	db := openDBConnection(t)
	defer db.Close()

	createTestCard(t)
}

func TestReadCard(t *testing.T) {
	db := openDBConnection(t)
	defer db.Close()

	createCardResponse := createTestCard(t)

	r := getRouter()
	req, err := http.NewRequest(http.MethodGet, "/v1/"+createCardResponse.Reference, nil)
	assert.Nil(t, err)

	req.Header.Set("Accept", "application/json")

	data := processRequest(r, req, t, http.StatusOK)

	response := CardResponse{}

	err = json.Unmarshal(data, &response)
	assert.Nil(t, err)

	assert.Equal(t, createCardResponse.Currency, response.Currency)
	assert.Equal(t, createCardResponse.Balance, response.Balance)
	assert.Equal(t, createCardResponse.ActivationDate.Unix(), response.ActivationDate.Unix())
	assert.Equal(t, createCardResponse.ExpireDate.Unix(), response.ExpireDate.Unix())
	assert.Equal(t, createCardResponse.Reference, response.Reference)
	assert.Equal(t, createCardResponse.CardNumber, response.CardNumber)
	assert.Equal(t, createCardResponse.Cvc, response.Cvc)
}

func TestDeleteCard(t *testing.T) {
	db := openDBConnection(t)
	defer db.Close()

	createCardResponse := createTestCard(t)

	r := getRouter()
	req, err := http.NewRequest(http.MethodDelete, "/v1/"+createCardResponse.Reference, nil)
	assert.Nil(t, err)

	processRequest(r, req, t, http.StatusNoContent)
}
