package card

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
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

func GetRouter() *gin.Engine {
	r := gin.Default()
	RegisterRoutes(r)

	return r
}

func CheckHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) error) {

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	if err := f(w); err != nil {
		t.Log(err)
		t.Fail()
	}
}

func OpenDBConnection(t *testing.T) *gorm.DB {
	var db *gorm.DB

	_, mock, err := sqlmock.NewWithDSN("sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}

	db, err = gorm.Open("sqlmock", "sqlmock_db_0")
	if err != nil {
		panic("Got an unexpected error.")
	}

	mock.ExpectBegin()
	mock.ExpectQuery("CREATE TABLE .**")
	common.DB = db

	AutoMigrate(db)

	mock.ExpectCommit()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	// To make http requests use our mock db
	return db
}

func TestCreateCard(t *testing.T) {
	db := OpenDBConnection(t)
	defer db.Close()

	r := GetRouter()
	activationDate := time.Now().Add(time.Hour * 100)
	expireDate := time.Now().Add(time.Hour * 150)
	currency := "EUR"

	payload := url.Values{}
	payload.Set("currency", currency)
	payload.Set("activationDate", activationDate.Format("2006-01-02"))
	payload.Set("expireDate", expireDate.Format("2006-01-02"))
	payload.Set("balance", "5000")

	req, err := http.NewRequest(http.MethodPost, "/v1/card/create", strings.NewReader(payload.Encode()))
	assert.Nil(t, err)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	CheckHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) error {
		assert.Equal(t, http.StatusCreated, w.Code)

		data, err := ioutil.ReadAll(w.Body)
		assert.Nil(t, err)

		response := CardResponse{}
		fmt.Println(string(data))

		err = json.Unmarshal(data, &response)
		assert.Nil(t, err)

		assert.Equal(t, currency, response.Currency)
		fmt.Println(response.Balance)
		assert.Equal(t, uint(5000), response.Balance)

		assert.NotNil(t, response.ActivationDate)
		assert.NotNil(t, response.ExpireDate)

		assert.NotNil(t, response.Reference)
		assert.NotNil(t, response.Cvc)
		assert.NotNil(t, response.CardNumber)

		return nil
	})
}
