package user

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

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type userResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Initialize gorm
//var dbGorm *gorm.DB

//func insertDummy(db *gorm.DB)  {
//	// Prepare dummy data
//	var newUser models.User
//	newUser.Name = "Test Login"
//	newUser.Email = "test1@test.com"
//	newUser.Password = "RAHASIA"
//
//	// user dummy data with model
//	customerModel := database.NewUserModel(db)
//	_, err := customerModel.Register(newUser)
//	if err != nil {
//		log.Fatalln(err)
//	}
//}

// TestController_RegisterUserController is unit testing for register controller
func TestController_RegisterUserController(t *testing.T) {
	//Initialize test cases
	var testCases = []struct {
		name            string
		reqBody         map[string]string
		expectCode      int
		responseStatus  string
		responseMessage string
	}{
		{
			name:            "test1",
			reqBody:         map[string]string{"name": "kuuga", "email": "kamen@rider.jp", "password": "kuuga99"},
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "test2",
			reqBody:         map[string]string{"name": "agito", "email": "", "password": ""},
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Register Failed",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Drop and create new table
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	// Initialize server
	e := echo.New()

	// Initialize user model
	modelUser := database.NewUserModel(db)

	// Initialize user controller
	controllerUser := NewController(modelUser)

	// Process all test cases
	for _, testCase := range testCases {
		body, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, controllerUser.RegisterUserController(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_LoginUserController(t *testing.T) {
	//Initialize test cases
	var testCases = []struct {
		name            string
		reqBody         map[string]string
		expectCode      int
		responseStatus  string
		responseMessage string
	}{
		{
			name:            "test1",
			reqBody:         map[string]string{"email": "test1@test.com", "password": "1234pass"},
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Login Success",
		},
		{
			name:            "test2",
			reqBody:         map[string]string{"email": "agito@kamen.jp", "password": "kuuga99"},
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Login Failed",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Drop and create new table
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	// Prepare dummy data
	//cost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))
	//password, _ := bcrypt.GenerateFromPassword([]byte("1234pass"), cost)
	newUser := models.User{
		Name:     "Test Login",
		Email:    "test1@test.com",
		Password: "1234pass",
	}
	registerModel := database.NewUserModel(db)
	_, err := registerModel.Register(newUser)
	if err != nil {
		fmt.Println(err)
	}

	// Initialize server
	e := echo.New()

	// Initialize user model
	modelUser := database.NewUserModel(db)

	// Initialize user controller
	controllerUser := NewController(modelUser)

	// Process all test cases
	for _, testCase := range testCases {
		body, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, controllerUser.LoginUserController(ctx)) {
			resBody := res.Body.String()

			fmt.Println("-----------------------------------------")
			fmt.Println(resBody)
			fmt.Println("-----------------------------------------")

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}
