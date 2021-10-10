package book

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"sewabuku/config"
	"sewabuku/database"
	"sewabuku/middlewares"
	"strings"
	"testing"

	m "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	e           = echo.New()
	db          = config.DBConnectTest()
	bookmodel   = database.NewBookModel(db)
	bookHandler = NewBookController(bookmodel)
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
		category   string
		author     string
		publisher  string
		expectCode int
		expectMsg  string
	}{
		{
			testName:   1,
			category:   "4",
			publisher:  "2",
			author:     "3",
			expectCode: http.StatusOK,
			expectMsg:  "success get books",
		}, {
			testName:   2,
			category:   "3",
			author:     "4",
			publisher:  "3",
			expectCode: http.StatusOK,
			expectMsg:  "no book found",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/books/search", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		q := req.URL.Query()
		q.Add("category", testCase.category)
		q.Add("author", testCase.author)
		q.Add("publisher", testCase.publisher)
		req.URL.RawQuery = q.Encode()
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := bookHandler.FilterAuthorCategoryPublisher(c)
		assert.NoError(t, err)
		responseMsg := res.Body.Bytes()
		var response bookResponse
		json.Unmarshal(responseMsg, &response)
		// fmt.Println(response, "kok kosong")
		// fmt.Println(res)
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
			expectMsg:  "success get all books",
		}, {
			testName:   2,
			expectCode: http.StatusOK,
			expectMsg:  "success get all books",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/all", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := bookHandler.GetAllBooks(c)
		assert.NoError(t, err)

		resBody := res.Body.Bytes()
		var response bookResponse
		json.Unmarshal(resBody, &response)

		if assert.NoError(t, bookHandler.GetAllBooks(c)) {

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.expectMsg, response.Message)
		}

	}

}

func TestSearch(t *testing.T) {
	testCases := []struct {
		testName   string
		keyword    string
		category   string
		publisher  string
		author     string
		expectCode int
		expectMsg  string
	}{
		{
			testName:   "test pertama",
			keyword:    "black",
			author:     "8",
			category:   "6",
			publisher:  "1",
			expectCode: http.StatusOK,
			expectMsg:  "success get books",
		}, {
			testName:   "test kedua",
			keyword:    "malam malam malam",
			expectCode: http.StatusOK,
			expectMsg:  "theres no book found",
		}, {
			testName:   "test3",
			keyword:    "Langit",
			author:     "6",
			publisher:  "3",
			expectCode: 200,
			expectMsg:  "success get books",
		},
	}
	for _, testCase := range testCases {
		fmt.Printf("test case =======================================================no %s", testCase.testName)
		req := httptest.NewRequest(http.MethodGet, "/books/search/", nil)
		res := httptest.NewRecorder()
		q := req.URL.Query()
		q.Add("category", testCase.category)
		q.Add("author", testCase.author)
		q.Add("publisher", testCase.publisher)

		req.URL.RawQuery = q.Encode()
		c := e.NewContext(req, res)
		c.SetPath("/:keyword")
		c.SetParamNames("keyword")
		c.SetParamValues(testCase.keyword)
		err := bookHandler.SearchAll(c)
		if err != nil {
			fmt.Println(err)
		}
		resBody := res.Body.Bytes()
		var response bookResponse

		err = json.Unmarshal(resBody, &response)
		if err != nil {
			return
		}

		fmt.Println(res)
		if assert.NoError(t, bookHandler.SearchAll(c)) {

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.expectMsg, response.Message)
		}

	}

}

func TestGetDetailsBook(t *testing.T) {
	testCases := []struct {
		name         string
		bookid       string
		expctedMSg   string
		expectedCode int
	}{
		{
			name:         "test1",
			bookid:       "1",
			expctedMSg:   "success get details book",
			expectedCode: 200,
		}, {
			name:         "test2",
			bookid:       "2",
			expctedMSg:   "success get details book",
			expectedCode: 200,
		},
	}

	for _, testCase := range testCases {
		fmt.Printf("test case =======================================================no==================no %s", testCase.name)
		req := httptest.NewRequest(http.MethodGet, "/books/details/", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		c.SetPath("/:id")
		c.SetParamNames("id")
		c.SetParamValues(testCase.bookid)
		err := bookHandler.GetDetailsBook(c)
		assert.NoError(t, err)
		resBody := res.Body.Bytes()
		var response bookResponse
		json.Unmarshal(resBody, &response)

		fmt.Println(response)

		if assert.NoError(t, bookHandler.GetDetailsBook(c)) {

			assert.Equal(t, testCase.expectedCode, res.Code)
			assert.Equal(t, testCase.expctedMSg, response.Message)
		}

	}
}

func TestGetPublisher(t *testing.T) {

	testCases := []struct {
		name         string
		expectedCode int
		ExpectMsg    string
	}{
		{
			name:         "test1",
			expectedCode: 200,
			ExpectMsg:    "success",
		}, {
			name:         "test2",
			expectedCode: 200,
			ExpectMsg:    "success",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/books/listpublisher", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := bookHandler.GetListPublisher(c)
		assert.NoError(t, err)
		if assert.NoError(t, bookHandler.GetListPublisher(c)) {
			assert.Equal(t, testCase.expectedCode, res.Code)
		}

	}
}

func TestGetAuthor(t *testing.T) {

	testCases := []struct {
		name         string
		expectedCode int
		ExpectMsg    string
	}{
		{
			name:         "test1",
			expectedCode: 200,
			ExpectMsg:    "success",
		}, {
			name:         "test2",
			expectedCode: 200,
			ExpectMsg:    "success",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/books/listauthor", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := bookHandler.GetListAuthor(c)
		assert.NoError(t, err)
		if assert.NoError(t, bookHandler.GetListAuthor(c)) {
			assert.Equal(t, testCase.expectedCode, res.Code)
		}

	}
}

func TestGetCategory(t *testing.T) {

	testCases := []struct {
		name         string
		expectedCode int
		ExpectMsg    string
	}{
		{
			name:         "test1",
			expectedCode: 200,
			ExpectMsg:    "success",
		}, {
			name:         "test2",
			expectedCode: 200,
			ExpectMsg:    "success",
		},
	}

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/books/listcategory", nil)
		res := httptest.NewRecorder()
		c := e.NewContext(req, res)
		err := bookHandler.GetListCategory(c)
		assert.NoError(t, err)
		if assert.NoError(t, bookHandler.GetListCategory(c)) {
			assert.Equal(t, testCase.expectedCode, res.Code)
		}

	}
}

func TestInsertBook(t *testing.T) {

	token1, _ := middlewares.CreateToken(2)
	token2, _ := middlewares.CreateToken(4)
	token3, _ := middlewares.CreateToken(5)

	testCases := []struct {
		name         string
		Req          map[string]interface{}
		token        string
		expectedMsg  string
		expectedCode int
	}{
		{
			name:         "test1",
			Req:          map[string]interface{}{"title": "mengejar mimpi", "publish_year": 2009, "description": "bismillah", "price": 100, "quantity": 1},
			token:        token1,
			expectedCode: 200,
			expectedMsg:  "successfully insert your new book",
		}, {
			name:         "test2",
			Req:          map[string]interface{}{"title": "meteor", "publish_year": 2009, "description": "bismillah", "price": 100, "quantity": 1, "publisher_id": 2},
			token:        token2,
			expectedCode: 200,
			expectedMsg:  "successfully insert your new book",
		}, {
			name:         "test3",
			Req:          map[string]interface{}{"title": "meteor", "description": "bismillah", "price": 100, "quantity": 1, "publisher_id": 2},
			token:        token3,
			expectedCode: 422,
			expectedMsg:  "please input name of your book and year of publishment of your book",
		},
	}

	for _, testCase := range testCases {
		reqBody, _ := json.Marshal(testCase.Req)
		req := httptest.NewRequest(http.MethodPost, "/books/newbook", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", testCase.token))
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)

		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(bookHandler.InsertBook)(ctx)) {
			var response bookResponse
			resBody := res.Body.Bytes()
			json.Unmarshal(resBody, &response)
			assert.Equal(t, testCase.expectedCode, res.Code)
			assert.Equal(t, testCase.expectedMsg, response.Message)

		}
	}
}

func TestCreateNewAuthor(t *testing.T) {
	token1, _ := middlewares.CreateToken(2)
	token2, _ := middlewares.CreateToken(4)
	token3, _ := middlewares.CreateToken(5)

	testCases := []struct {
		nameTest     string
		req          string
		token        string
		expectedMsg  string
		expectedCode int
	}{
		{
			nameTest:     "test pertama",
			req:          "sarah wijayantia lestaluhu",
			token:        token1,
			expectedCode: 200,
			expectedMsg:  "successfully create new author",
		}, {
			nameTest:     "test 2",
			req:          "tere li*ye",
			token:        token2,
			expectedCode: 422,
			expectedMsg:  "cannot create new author if there are special characters",
		}, {
			nameTest:     "test3",
			req:          "tere liye",
			token:        token3,
			expectedCode: 422,
			expectedMsg:  "cannot input name author with same name which already exist",
		},
	}

	for _, test := range testCases {
		f := make(url.Values)
		f.Set("name", test.req)
		req := httptest.NewRequest(http.MethodPost, "/books/newauthor", strings.NewReader(f.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)

		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(bookHandler.CreateNewAuthor)(ctx)) {
			fmt.Println(res)
			var response bookResponse
			resBody := res.Body.Bytes()
			json.Unmarshal(resBody, &response)
			assert.Equal(t, test.expectedCode, res.Code)
			assert.Equal(t, test.expectedMsg, response.Message)

		}

	}

}

func TestCreateNewPublisher(t *testing.T) {

	token1, _ := middlewares.CreateToken(2)
	testCases := []struct {
		nameTest    string
		Request     string
		token       string
		expectdCode int
		ExpectedMsg string
	}{
		{
			nameTest:    "test1",
			Request:     "gramedia",
			token:       token1,
			expectdCode: 422,
			ExpectedMsg: "cannot input name author with same name which already exist",
		}, {
			nameTest:    "test2",
			Request:     "baruajacobayaa",
			token:       token1,
			expectdCode: 200,
			ExpectedMsg: "successfully create new publisher",
		}, {
			nameTest:    "test3",
			Request:     "*hamar%^&",
			token:       token1,
			expectdCode: 422,
			ExpectedMsg: "cannot create new author if there are special characters",
		},
	}

	for _, test := range testCases {
		f := url.Values{}
		f.Set("name", test.Request)
		req := httptest.NewRequest(http.MethodPost, "/books/newpublisher", strings.NewReader(f.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(bookHandler.CreateNewPublisher)(ctx)) {

			responseBody := res.Body.Bytes()
			var response bookResponse

			json.Unmarshal(responseBody, &response)
			assert.Equal(t, test.expectdCode, res.Code)
			assert.Equal(t, test.ExpectedMsg, response.Message)

		}
	}
}

// func TestUpdatePhotoBook(t *testing.T) {

// 	token1, _ := middlewares.CreateToken(1)
// 	token2, _ := middlewares.CreateToken(2)

// 	testCases := []struct {
// 		nameTest     string
// 		token        string
// 		bookid       string
// 		ReqFile      string
// 		expectdMsg   string
// 		expectedCode int
// 	}{
// 		{
// 			nameTest:     "test1",
// 			token:        token1,
// 			ReqFile:      "/home/rjandoni/Desktop/draf-sewabuku/image/4,1.png",
// 			bookid:       "5",
// 			expectdMsg:   "successfully update books photo",
// 			expectedCode: 200,
// 		}, {
// 			nameTest:     "test2",
// 			token:        token2,
// 			bookid:       "1",
// 			ReqFile:      "/home/rjandoni/Downloads/bwa-storegg-server.zip",
// 			expectdMsg:   "failed",
// 			expectedCode: 422,
// 		},
// 	}

// 	for _, test := range testCases {
// 		fmt.Printf("================================   %s   =========================", test.nameTest)
// 		pr, pw := io.Pipe()

// 		writer := multipart.NewWriter(pw)

// 		go func() {
// 			defer writer.Close()

// 			part, err := writer.CreateFormFile("file", test.ReqFile)
// 			assert.NoError(t, err)

// 			err = png.Encode(part, nil)
// 			assert.NoError(t, err)

// 		}()

// 		req := httptest.NewRequest(http.MethodPut, "/books/bookphoto/", pr)
// 		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
// 		req.Header.Add(echo.HeaderContentType, writer.FormDataContentType())

// 		res := httptest.NewRecorder()

// 		ctx := e.NewContext(req, res)
// 		ctx.Param("/:id")
// 		ctx.SetParamNames("id")
// 		ctx.SetParamValues(test.bookid)

// 		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(bookHandler.UpdatePhotoBook)(ctx)) {
// 			fmt.Println(res)
// 			assert.Equal(t, test.expectedCode, res.Code)
// 		}
// 	}

// }
