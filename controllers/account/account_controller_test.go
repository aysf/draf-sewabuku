package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sewabuku/config"
	"sewabuku/database"
	"sewabuku/middlewares"
	"testing"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

var (
	e              = echo.New()
	db             = config.DBConnectTest()
	accountModel   = database.NewAccountModel(db)
	accountHandler = NewController(accountModel)
)

func TestTopUpWithdraw(t *testing.T) {
	token1, err := middlewares.CreateToken(1)
	assert.NoError(t, err)

	token2, err := middlewares.CreateToken(2)
	assert.NoError(t, err)

	testCases := []struct {
		testName   string
		querry     string
		token      string
		request    map[string]int
		expectCode int
		expectMsg  string
	}{
		{
			testName:   "test1",
			querry:     "1",
			token:      token1,
			request:    map[string]int{"amount": 3000000},
			expectCode: 200,
			expectMsg:  "success",
		}, {
			testName:   "test2",
			querry:     "1",
			token:      token2,
			request:    map[string]int{"amount": 20000},
			expectCode: 200,
			expectMsg:  "success",
		}, {
			testName:   "test2",
			querry:     "2",
			token:      token1,
			request:    map[string]int{"amount": 10000},
			expectCode: 200,
			expectMsg:  "success",
		},
	}

	for _, test := range testCases {
		reqBody, err := json.Marshal(test.request)
		assert.Nil(t, err)

		req := httptest.NewRequest(http.MethodPost, "/account/transaction", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		q := req.URL.Query()
		q.Add("code", test.querry)
		req.URL.RawQuery = q.Encode()

		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		if assert.Nil(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(accountHandler.TopupWithdraw)(ctx)) {
			assert.Equal(t, test.expectCode, res.Code)
			fmt.Println(res)

		}

	}

}

func TestDepositTransfer(t *testing.T) {
	token1, err := middlewares.CreateToken(1)
	assert.NoError(t, err)

	token2, err := middlewares.CreateToken(5)
	assert.NoError(t, err)

	token3, err := middlewares.CreateToken(1)
	assert.NoError(t, err)

	testCases := []struct {
		testName   string
		token      string
		request    map[string]interface{}
		expectCode int
		expectMsg  string
	}{
		{
			testName:   "test1",
			token:      token1,
			request:    map[string]interface{}{"Balance": 40000},
			expectCode: 200,
			expectMsg:  "success",
		}, {
			testName:   "test1",
			token:      token2,
			request:    map[string]interface{}{"Balance": 1000},
			expectCode: 200,
			expectMsg:  "success",
		}, {
			testName:   "test1",
			token:      token3,
			request:    map[string]interface{}{"Balance": 1000},
			expectCode: 200,
			expectMsg:  "success",
		},
	}

	for _, test := range testCases {
		reqBody, err := json.Marshal(test.request)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/account/deposit", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)
		if assert.Nil(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(accountHandler.DepositTransfer)(ctx)) {
			fmt.Println(res)
			assert.Equal(t, test.expectCode, res.Code)
		}
	}
}
