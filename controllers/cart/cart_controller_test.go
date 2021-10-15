package cart

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

	m "github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type cartResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// func insertDummyData(db *gorm.DB) {
// 	pass, _ := bcrypt.GenerateFromPassword([]byte("123"), 14)
// 	passStr := string(pass)
// 	var user = []models.User{
// 		{Name: "Ami", Email: "ami@mail.com", Password: passStr, Address: "jakarta"},
// 		{Name: "Baiq", Email: "baiq@mail.com", Password: passStr, Address: "depok"},
// 	}
// 	var entry = []models.Entry{
// 		{AccountID: "a-1", Amount: 25000, CreatedAt: time.Now()},
// 		{AccountID: "a-2", Amount: 0, CreatedAt: time.Now()},
// 	}
// 	var bookData = []models.BookData{
// 		{Title: "Rich Dad Poor Dad", UserID: 1, Quantity: 1, Photo: "default.jpg", AuthorID: 1, PublisherID: 1, CategoryID: 1, PublishYear: 1997, Price: 5000},
// 		{Title: "Kambing Jantan", UserID: 2, Quantity: 2, Photo: "default.jpg", AuthorID: 1, PublisherID: 1, CategoryID: 1, PublishYear: 1997, Price: 100},
// 	}
// 	db.Create(&user)
// 	db.Create(&entry)
// 	db.Create(&bookData)
// }

var (
	e           = echo.New()
	db          = config.DBConnectTest()
	cartmodel   = database.NewCartModel(db)
	cartHandler = NewCartController(cartmodel)
)

func TestController_RentBook(t *testing.T) {

	token1, err := middlewares.CreateToken(1)
	assert.NoError(t, err)
	token2, err := middlewares.CreateToken(2)
	assert.NoError(t, err)

	var testCases = []struct {
		name            string
		token           string
		reqBody         map[string]interface{}
		expectCode      int
		responseStatus  string
		responseMessage string
	}{
		{
			name:            "case 1: user 1 borrow to user 2 - rent available books",
			reqBody:         map[string]interface{}{"book_data_id": 1, "date_loan": "2021-09-01T00:00:00+07:00", "date_due": "2021-09-01T00:00:00+07:00", "date_return": "2021-09-01T00:00:00+07:00"},
			token:           token1,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 2: user 2 borrow to user 1 - rent books when does not have enough balance",
			reqBody:         map[string]interface{}{"book_data_id": 2, "date_loan": "2021-09-01T00:00:00+07:00", "date_due": "2021-09-05T00:00:00+07:00", "date_return": "2021-09-05T00:00:00+07:00"},
			token:           token2,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 3: user 1 borrow again to user 2 - rent the book id that already rented before",
			reqBody:         map[string]interface{}{"book_data_id": 2, "date_loan": "2021-09-01T00:00:00+07:00", "date_due": "2021-09-05T00:00:00+07:00", "date_return": "2021-09-05T00:00:00+07:00"},
			token:           token1,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 4: user 1 borrow book id 1 - rent book for user itself",
			reqBody:         map[string]interface{}{"book_data_id": 1, "date_loan": "2021-09-01T00:00:00+07:00", "date_due": "2021-09-05T00:00:00+07:00", "date_return": "2021-09-05T00:00:00+07:00"},
			token:           token1,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
	}

	// // Initialize database connection
	// db := config.DBConnectTest()

	// // Drop and create new table
	// db.Migrator().DropTable(
	// 	&models.Account{},
	// 	&models.User{},
	// 	&models.Entry{},
	// 	&models.BookData{},
	// )
	// db.AutoMigrate(
	// 	&models.Account{},
	// 	&models.User{},
	// 	&models.Entry{},
	// 	&models.BookData{},
	// )

	// // Prepare dummy data
	// insertDummyData(db)

	// // // Initialize server
	// e := echo.New()

	// // Initialize user model
	// modelCart := database.NewCartModel(db)

	// // Initialize user controller
	// controllerCart := NewCartController(modelCart)

	//loop
	for _, tc := range testCases {
		body, err := json.Marshal(tc.reqBody)
		assert.NoError(t, err)
		req := httptest.NewRequest(echo.POST, "/carts/rent", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", tc.token))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(cartHandler.RentBook)(ctx)) {
			resBody := res.Body.Bytes()

			var response cartResponse
			err := json.Unmarshal(resBody, &response)
			assert.Nil(t, err)
			fmt.Println(res)

			assert.Equal(t, tc.expectCode, res.Code)
			assert.Equal(t, tc.responseStatus, response.Status)
			assert.Equal(t, tc.responseMessage, response.Message)

		}
	}

	// temp
}

func TestGetListBook(t *testing.T) {
	token1, err := middlewares.CreateToken(5)
	assert.NoError(t, err)
	token2, err := middlewares.CreateToken(6)
	assert.NoError(t, err)
	token3, err := middlewares.CreateToken(4)
	assert.Nil(t, err)

	testCases := []struct {
		testname   string
		token      string
		expectCode int
		expectMsg  string
	}{
		{
			testname:   "TEST 1",
			token:      token1,
			expectCode: 200,
			expectMsg:  "Success get book borrowing list",
		}, {
			testname:   "test 2",
			token:      token2,
			expectCode: 200,
			expectMsg:  "Success get book borrowing list",
		}, {
			testname:   "test 3",
			token:      token3,
			expectCode: 200,
			expectMsg:  "Success get book borrowing list",
		},
	}

	for _, test := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)

		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(cartHandler.ListBook)(ctx)) {
			fmt.Println(res)
			var response cartResponse
			err := json.Unmarshal(res.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, test.expectCode, res.Code)
			assert.Equal(t, test.expectMsg, response.Message)
		}
	}
}

func TestExtenntDateDue(t *testing.T) {
	token, err := middlewares.CreateToken(5)
	assert.NoError(t, err)
	token1, err := middlewares.CreateToken(6)
	assert.Nil(t, err)
	testCases := []struct {
		testname   string
		token      string
		request    map[string]interface{}
		expectcode int
		expectmsg  string
	}{
		{
			testname:   "test1",
			token:      token,
			request:    map[string]interface{}{"book_data_id": 5, "date_due": "2021-09-05T00:00:00+07:00"},
			expectcode: 200,
			expectmsg:  "succeess",
		}, {
			testname:   "test2",
			token:      token1,
			request:    map[string]interface{}{"book_data_id": 5, "date_due": "2021-09-05T00:00:00+07:00"},
			expectcode: 400,
			expectmsg:  "Fail to extend date due",
		},
	}

	for _, test := range testCases {
		reqbody, err := json.Marshal(test.request)
		assert.NoError(t, err)
		req := httptest.NewRequest(http.MethodPut, "/carts/extend", bytes.NewBuffer(reqbody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(cartHandler.ExtendDateDue)(ctx)) {
			fmt.Println(res)
			var response cartResponse
			err := json.Unmarshal(res.Body.Bytes(), &response)
			assert.NoError(t, err)

			assert.Equal(t, test.expectcode, res.Code)
			assert.Equal(t, test.expectmsg, response.Message)
		}
	}
}

func TestReturnBook(t *testing.T) {
	token, err := middlewares.CreateToken(6)
	assert.NoError(t, err)

	token2, err := middlewares.CreateToken(3)
	assert.NoError(t, err)

	TestCases := []struct {
		testname     string
		token        string
		request      map[string]interface{}
		expextedcode int
	}{
		{
			testname:     "test1",
			token:        token,
			request:      map[string]interface{}{"book_data_id": 3, "date_return": "2021-09-05T00:00:00+07:00"},
			expextedcode: http.StatusBadRequest,
		}, {
			testname:     "test2",
			token:        token2,
			request:      map[string]interface{}{"book_data_id": 4, "date_return": "2021-09-05T00:00:00+07:00"},
			expextedcode: http.StatusBadRequest,
		},
	}

	for _, test := range TestCases {
		reqBody, err := json.Marshal(test.request)
		assert.NoError(t, err)

		req := httptest.NewRequest(http.MethodPut, "/carts/return", bytes.NewBuffer(reqBody))
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
		res := httptest.NewRecorder()

		ctx := e.NewContext(req, res)

		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(cartHandler.ReturnBook)(ctx)) {
			fmt.Println(res)
			assert.Equal(t, test.expextedcode, res.Code)
		}

	}
}
