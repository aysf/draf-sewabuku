package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"sewabuku/config"
	"sewabuku/database"
	"sewabuku/models"
	"testing"
)

type userResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// TestController_RegisterUserController is unit test for RegisterUserController
func TestRegisterUserController(t *testing.T) {
	// Initialize test cases
	var testCases = []struct {
		name            string
		reqBody         map[string]string
		expectCode      int
		responseStatus  string
		responseMessage string
	}{
		{
			name:            "test1",
			reqBody:         map[string]string{"name": ""},
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Name, Email, or Password cannot Null",
		},
		{
			name:            "test2",
			reqBody:         map[string]string{"name": "Kuuga", "email": "kamen@rider.jp", "password": "kuuga99"},
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "test3",
			reqBody:         map[string]string{"name": "Kuuga2", "email": "kamen2@rider.jp", "password": "kuuga99"},
			expectCode:      http.StatusInternalServerError,
			responseStatus:  "fail",
			responseMessage: "Register Fail",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Drop and create new table
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})

	e := echo.New()

	// Setting request
	for i, testCase := range testCases {
		fmt.Println("======================================")
		fmt.Println("ini test ke", i)
		// Create new repo and controller user
		modelUser := database.NewUserModel(db)
		controllerUser := NewController(modelUser)
		body, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)

		if testCase.name == "test3" {
			db.Migrator().DropTable(&models.User{})
		}

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

//type userResponse struct {
//	Status  string      `json:"status"`
//	Message string      `json:"message"`
//	Data    interface{} `json:"data"`
//}
//
//func TestMain(m *testing.M) {
//	setup()
//	os.Exit(m.Run())
//}
//
//func setup() {
//	// create database connection
//	db := config.DBConnectTest()
//
//	// cleaning data before testing
//	db.Migrator().DropTable(&models.User{})
//	db.AutoMigrate(&models.User{})
//
//	// preparate dummy data
//	var newUser models.User
//	newUser.Name = "Name Test B"
//	newUser.Email = "testb@alterra.id"
//	newUser.Password = "password123"
//
//	// user dummy data with model
//	userModel := database.NewUserModel(db)
//	_, err := userModel.Register(newUser)
//	if err != nil {
//		fmt.Println(err)
//	}
//}

//func TestRegisterUserController(t *testing.T) {
//	// create database connection and create controller
//	db := config.DBConnectTest()
//	userModel := database.NewUserModel(db)
//	userController := NewController(userModel)
//
//	// input controller
//	reqBody, _ := json.Marshal(map[string]string{
//		"name":     "Name Test",
//		"email":    "test@alterra.id",
//		"password": "test123",
//	})
//
//	// setting controller
//	e := echo.New()
//	req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(reqBody))
//	res := httptest.NewRecorder()
//	req.Header.Set("Content-Type", "application/json")
//	context := e.NewContext(req, res)
//	context.SetPath("/users/register")
//
//	userController.RegisterUserController(context)
//
//	// build struct response
//	type Response struct {
//		Code    int    `json:"code"`
//		Message string `json:"message"`
//	}
//	var response Response
//	resBody := res.Body.String()
//	json.Unmarshal([]byte(resBody), &response)
//
//	// testing stuff
//	t.Run("POST /users/register", func(t *testing.T) {
//		assert.Equal(t, 200, res.Code)
//		assert.Equal(t, "Register Success", response.Message)
//	})
//}

//// TestController_RegisterUserController is unit test for RegisterUserController
//func TestController_RegisterUserController(t *testing.T) {
//	// Initialize test cases
//	var testCases = []struct {
//		reqBody         map[string]string
//		expectCode      int
//		responseStatus  string
//		responseMessage string
//	}{
//		{
//			reqBody:         map[string]string{"name": ""},
//			expectCode:      http.StatusInternalServerError,
//			responseStatus:  "fail",
//			responseMessage: "Register Success",
//		},
//		{
//			reqBody:         map[string]string{"name": "Kuuga", "email": "kamen@rider.jp", "password": "kuuga99"},
//			expectCode:      http.StatusOK,
//			responseStatus:  "success",
//			responseMessage: "Register Success",
//		},
//	}
//
//	// Initialize database connection
//	db := config.DBConnectTest()
//
//	// Drop and create new table
//	//db.Migrator().DropTable(&models.User{})
//	//db.AutoMigrate(&models.User{})
//
//
//	e := echo.New()
//
//	// Setting request
//	for _, testCase := range testCases {
//		// Create new repo and controller user
//		modelUser := database.NewUserModel(db)
//		controllerUser := NewController(modelUser)
//		body, _ := json.Marshal(testCase.reqBody)
//		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
//		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//		res := httptest.NewRecorder()
//		ctx := e.NewContext(req, res)
//
//		if assert.NoError(t, controllerUser.RegisterUserController(ctx)) {
//			resBody := res.Body.String()
//
//			var response userResponse
//			json.Unmarshal([]byte(resBody), &response)
//
//			t.Run("POST /users/register", func(t *testing.T) {
//				//		assert.Equal(t, 200, res.Code)
//				//		assert.Equal(t, "Register Success", response.Message)
//				assert.Equal(t, testCase.expectCode, res.Code)
//				assert.Equal(t, testCase.responseStatus, response.Status)
//				assert.Equal(t, testCase.responseMessage, response.Message)
//			})
//		}
//	}
//}
