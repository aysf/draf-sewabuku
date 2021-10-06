package book

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sewabuku/config"
	"sewabuku/database"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// import "testing"
var (
	e           = echo.New()
	db          = config.DBConnect()
	bookmodel   = database.NewBookModel(db)
	bookHandler = NewBookController(bookmodel)
)

func TestGetByKeyWord(t *testing.T) {
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

	for _, testCase := range testCases {
		reqBody := fmt.Sprintf("%d", testCase.category)
		req := httptest.NewRequest(http.MethodGet, "/books/search", strings.NewReader(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		bookHandler.FilterAuthorCategoryPublisher(c)
		if assert.NoError(t, bookHandler.FilterAuthorCategoryPublisher(c)) {

			assert.Equal(t, testCase.expectCode, res.Code)
		}

	}

}
