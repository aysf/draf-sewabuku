package user

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
	"sewabuku/models"
	"testing"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

type userResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

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
			reqBody:         map[string]string{"name": "kuuga", "email": "kamen@rider.jp", "password": "kuuga99", "address": "japan"},
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "test2",
			reqBody:         map[string]string{"name": "agito", "email": "", "password": "", "address": ""},
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Check Your Input",
		},
		{
			name:            "test3",
			reqBody:         map[string]string{"name": "agito", "email": "", "password": "", "address": ""},
			expectCode:      http.StatusInternalServerError,
			responseStatus:  "error",
			responseMessage: "Register Failed",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Drop and create new table
	db.Migrator().DropTable(
		&models.Account{},
		&models.User{},
	)
	db.AutoMigrate(
		&models.Account{},
		&models.User{},
	)

	// Initialize server
	e := echo.New()

	// Initialize user model
	modelUser := database.NewUserModel(db)

	// Initialize user controller
	controllerUser := NewController(modelUser)

	// Process all test cases
	for _, testCase := range testCases {
		fmt.Println("------------>", testCase.name)
		if testCase.name == "test3" {
			db.Migrator().DropTable(
				&models.User{},
				&models.Account{},
			)
		}
		body, _ := json.Marshal(testCase.reqBody)
		fmt.Println("==============request")
		fmt.Println(testCase.reqBody)
		req := httptest.NewRequest(echo.POST, "/users/register", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, controllerUser.RegisterUserController(ctx)) {
			resBody := res.Body.String()

			fmt.Println("==============response")
			fmt.Println(resBody)

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
		req := httptest.NewRequest(echo.POST, "/users/login", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, controllerUser.LoginUserController(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_GetUserProfileController(t *testing.T) {
	//Initialize test cases
	var testCases = []struct {
		name            string
		expectCode      int
		responseStatus  string
		responseMessage string
	}{
		{
			name:            "test1",
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Success Get User Profile",
		},
		{
			name:            "test2",
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Fail to Get User Profile",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Drop and create new table
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	// Prepare dummy data
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

	// Create token
	token, _ := middlewares.CreateToken(1)

	// Initialize server
	e := echo.New()

	// Initialize user model
	modelUser := database.NewUserModel(db)

	// Initialize user controller
	controllerUser := NewController(modelUser)

	// Process all test cases
	for _, testCase := range testCases {
		req := httptest.NewRequest(echo.GET, "/users/profile", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		if testCase.name == "test2" {
			db.Migrator().DropTable(&models.User{})
		}
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.GetUserProfileController)(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_UpdateUserProfileController(t *testing.T) {
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
			reqBody:         map[string]string{"address": "singapore"},
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Success Update Profile",
		},
		{
			name:            "test2",
			reqBody:         map[string]string{"email": "ubahemail@com"},
			expectCode:      http.StatusInternalServerError,
			responseStatus:  "error",
			responseMessage: "Check Your Input",
		},
		{
			name:            "test3",
			reqBody:         map[string]string{"name": "Yuri"},
			expectCode:      http.StatusInternalServerError,
			responseStatus:  "error",
			responseMessage: "Check Your Input",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Drop and create new table
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	// Prepare dummy data
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

	// Create token
	token, _ := middlewares.CreateToken(1)

	// Initialize server
	e := echo.New()

	// Initialize user model
	modelUser := database.NewUserModel(db)

	// Initialize user controller
	controllerUser := NewController(modelUser)

	// Process all test cases
	for _, testCase := range testCases {
		body, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(echo.GET, "/users/profile", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		if testCase.name == "test2" {
			db.Migrator().DropTable(&models.User{})
		}
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.GetUserProfileController)(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}
