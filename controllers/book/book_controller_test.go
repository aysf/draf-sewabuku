package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sewabuku/config"
	"sewabuku/database"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// import "testing"

type bookResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func TestFilterring(t *testing.T) {
	testCases := []struct {
		testName   int
		category   int
		author     int
		publisher  int
		expectCode int
		expectMsg  string
	}{
		{
			testName:   1,
			category:   2,
			publisher:  1,
			author:     3,
			expectCode: http.StatusOK,
			expectMsg:  "success",
		}, {
			testName:   2,
			category:   3,
			author:     2,
			publisher:  2,
			expectCode: http.StatusOK,
			expectMsg:  "success",
		},
	}

	e := echo.New()
	db := config.DBConnectTest()
	bookmodel := database.NewBookModel(db)
	bookHandler := NewBookController(bookmodel)
	for _, testCase := range testCases {
		reqBody := fmt.Sprintf("?publisher=%d&?category=%d", testCase.publisher, testCase.category)
		reqBodyJson, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodGet, "/books/s", bytes.NewBuffer(reqBodyJson))
		// req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		responseMsg := res.Body.String()
		c := e.NewContext(req, res)
		err := bookHandler.FilterAuthorCategoryPublisher(c)
		if err != nil {
			fmt.Println(err)
		}
		var response bookResponse
		json.Unmarshal([]byte(responseMsg), &response)
		fmt.Println(res)
		if assert.NoError(t, bookHandler.FilterAuthorCategoryPublisher(c)) {

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.expectMsg, response.Message)
		}

	}

}

func TestGetALL(t *testing.T) {
	testCases := []struct {
		testName int

		expectCode int
		expectMsg  string
	}{
		{
			testName:   1,
			expectCode: http.StatusOK,
			expectMsg:  "success",
		}, {
			testName:   2,
			expectCode: http.StatusOK,
			expectMsg:  "success",
		},
	}

	e := echo.New()
	db := config.DBConnectTest()
	bookmodel := database.NewBookModel(db)
	bookHandler := NewBookController(bookmodel)
	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/all", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := bookHandler.GetAllBooks(c)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)

		if assert.NoError(t, bookHandler.GetAllBooks(c)) {

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.expectMsg, res.Result().Body)
		}

	}

}
