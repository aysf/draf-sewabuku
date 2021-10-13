package cart

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sewabuku/config"
	"sewabuku/database"
	"sewabuku/models"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type cartResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func insertDummyData(db *gorm.DB) {
	pass, _ := bcrypt.GenerateFromPassword([]byte("123"), 14)
	passStr := string(pass)
	var user = []models.User{
		{Name: "Ami", Email: "ami@mail.com", Password: passStr, Address: "jakarta"},
		{Name: "Baiq", Email: "baiq@mail.com", Password: passStr, Address: "depok"},
	}
	var entry = []models.Entry{
		{AccountID: "a-1", Amount: 25000, CreatedAt: time.Now()},
		{AccountID: "a-2", Amount: 0, CreatedAt: time.Now()},
	}
	var bookData = []models.BookData{
		{Title: "Rich Dad Poor Dad", UserID: 1, Quantity: 1, Photo: "default.jpg", AuthorID: 1, PublisherID: 1, CategoryID: 1, PublishYear: 1997, Price: 5000},
		{Title: "Kambing Jantan", UserID: 2, Quantity: 2, Photo: "default.jpg", AuthorID: 1, PublisherID: 1, CategoryID: 1, PublishYear: 1997, Price: 100},
	}
	db.Create(&user)
	db.Create(&entry)
	db.Create(&bookData)
}

func TestController_RentBook(t *testing.T) {
	var testCases = []struct {
		name            string
		reqBody         string
		expectCode      int
		responseStatus  string
		responseMessage string
	}{
		{
			name:            "case 1: user 1 borrow to user 2 - rent available books",
			reqBody:         `{"book_data_id": 1, "date_loan":"2021-09-01T00:00:00+07:00", "date_due":"2021-09-05T00:00:00+07:00"}`,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 2: user 2 borrow to user 1 - rent books when does not have enough balance",
			reqBody:         ``,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 3: user 1 borrow again to user 2 - rent the book id that already rented before",
			reqBody:         ``,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 4: user 1 borrow book id 1 - rent book for user itself",
			reqBody:         ``,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Drop and create new table
	db.Migrator().DropTable(
		&models.Account{},
		&models.User{},
		&models.Entry{},
		&models.BookData{},
	)
	db.AutoMigrate(
		&models.Account{},
		&models.User{},
		&models.Entry{},
		&models.BookData{},
	)

	// Prepare dummy data
	insertDummyData(db)

	// Initialize server
	e := echo.New()

	// Initialize user model
	modelCart := database.NewCartModel(db)

	// Initialize user controller
	controllerCart := NewCartController(modelCart)

	//loop
	for _, tc := range testCases {
		body, _ := json.Marshal(tc.reqBody)
		req := httptest.NewRequest(echo.POST, "/users/login", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, controllerCart.RentBook(ctx)) {
			resBody := res.Body.String()

			var response cartResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, tc.expectCode, res.Code)
			assert.Equal(t, tc.responseStatus, response.Status)
			assert.Equal(t, tc.responseMessage, response.Message)

		}
	}

	// temp
	fmt.Print(db, testCases)
}
