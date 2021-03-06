package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
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

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

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
			reqBody:         map[string]string{"name": "kuuga", "organization_name": "Kamen Rider", "email": "kamen@rider.jp", "password": "kuuga99", "address": "japan"},
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
			reqBody:         map[string]string{"name": "ryuuki", "organization_name": "Kamen Rider", "email": "ryuki@rider.jp", "password": "ryukimaru", "address": "japan"},
			expectCode:      http.StatusInternalServerError,
			responseStatus:  "error",
			responseMessage: "Register Failed",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Initialize server
	e := echo.New()

	// Add validator module
	e.Validator = &CustomValidator{validator: validator.New()}

	// Initialize user model
	modelUser := database.NewUserModel(db)

	// Initialize user controller
	controllerUser := NewController(modelUser)

	// Process all test cases
	for _, testCase := range testCases {
		if testCase.name == "test3" {
			db.Migrator().DropTable(
				&models.User{},
				&models.Account{},
			)
		}
		body, _ := json.Marshal(testCase.reqBody)
		fmt.Println(testCase.reqBody)
		req := httptest.NewRequest(echo.POST, "/users/register", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, controllerUser.RegisterUserController(ctx)) {
			resBody := res.Body.String()

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
			fmt.Println(res)

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
			reqBody:         map[string]string{"name": "Yuri"},
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Fail to Update User Profile",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

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
		req := httptest.NewRequest(echo.POST, "/users/profile", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		if testCase.name == "test2" {
			db.Migrator().DropTable(&models.User{})
		}
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.UpdateUserProfileController)(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_UpdatePasswordController(t *testing.T) {
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
			reqBody:         map[string]string{"password": "newPass"},
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Success Change Password",
		},
		{
			name:            "test2",
			reqBody:         map[string]string{"password": ""},
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Fail to Change Password",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

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
		req := httptest.NewRequest(echo.POST, "/users/change-password", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		if testCase.name == "test2" {
			db.Migrator().DropTable(&models.User{})
		}
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.UpdatePasswordController)(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_LogoutUserController(t *testing.T) {
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
			responseMessage: "Logout Success",
		},
		{
			name:            "test2",
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Logout Failed",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

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
		req := httptest.NewRequest(echo.GET, "/users/logout", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		if testCase.name == "test2" {
			db.Migrator().DropTable(&models.User{})
		}
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.LogoutUserController)(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_GetBorrowedController(t *testing.T) {
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
			responseMessage: "Success Get Borrowed Book",
		},
		{
			name:            "test2",
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Fail to Get Borrowed Book",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

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
		req := httptest.NewRequest(echo.GET, "/users/cart", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		if testCase.name == "test2" {
			db.Migrator().DropTable(&models.User{})
		}
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		ctx.QueryParam("complete")
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.GetBorrowedController)(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_GetLentController(t *testing.T) {
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
			responseMessage: "Success Get Lent Book",
		},
		{
			name:            "test2",
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Fail to Get Lent Book",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

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
		req := httptest.NewRequest(echo.GET, "/users/books", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		if testCase.name == "test2" {
			db.Migrator().DropTable(&models.User{})
		}
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		ctx.QueryParam("complete")
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.GetLentController)(ctx)) {
			resBody := res.Body.String()

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_InsertRatingBookController(t *testing.T) {
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
			reqBody:         map[string]string{"rate_book": "5", "desc_rate_book": "Bukunya bagus bingit"},
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Success Give Book Rating",
		},
		{
			name:            "test2",
			reqBody:         map[string]string{"rate_book": "2", "desc_rate_book": ""},
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Fail to Give Book Rating",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Initialize server
	e := echo.New()

	// Add validator module
	e.Validator = &CustomValidator{validator: validator.New()}

	// Initialize user model
	modelUser := database.NewUserModel(db)

	// Initialize user controller
	controllerUser := NewController(modelUser)

	// Create token
	token, _ := middlewares.CreateToken(1)

	// Process all test cases
	for _, testCase := range testCases {
		if testCase.name == "test2" {
			db.Migrator().DropTable(
				&models.Rating{},
			)
		}
		body, _ := json.Marshal(testCase.reqBody)
		fmt.Println(testCase.reqBody)
		req := httptest.NewRequest(echo.POST, "/users/book-rating", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		ctx.SetPath("/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.InsertRatingBookController)(ctx)) {
			resBody := res.Body.String()

			fmt.Println(resBody)

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}

func TestController_InsertRatingBorrowerController(t *testing.T) {
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
			reqBody:         map[string]string{"rate_borrower": "5", "desc_rate_borrower": "Bukunya bagus bingit"},
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Success Give Book Rating",
		},
		{
			name:            "test2",
			reqBody:         map[string]string{"rate_borrower": "2", "desc_rate_borrower": ""},
			expectCode:      http.StatusBadRequest,
			responseStatus:  "fail",
			responseMessage: "Fail to Give Book Rating",
		},
	}

	// Initialize database connection
	db := config.DBConnectTest()

	// Initialize server
	e := echo.New()

	// Add validator module
	e.Validator = &CustomValidator{validator: validator.New()}

	// Initialize user model
	modelUser := database.NewUserModel(db)

	// Initialize user controller
	controllerUser := NewController(modelUser)

	// Create token
	token, _ := middlewares.CreateToken(1)

	// Process all test cases
	for _, testCase := range testCases {
		if testCase.name == "test2" {
			db.Migrator().DropTable(
				&models.Rating{},
			)
		}
		body, _ := json.Marshal(testCase.reqBody)
		fmt.Println(testCase.reqBody)
		req := httptest.NewRequest(echo.POST, "/users/borrower-rating", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %v", token))
		res := httptest.NewRecorder()
		ctx := e.NewContext(req, res)
		ctx.SetPath("/:id")
		ctx.SetParamNames("id")
		ctx.SetParamValues("1")
		if assert.NoError(t, m.JWT([]byte(os.Getenv("SECRET_KEY")))(controllerUser.InsertRatingBorrowerController)(ctx)) {
			resBody := res.Body.String()

			fmt.Println(resBody)

			var response userResponse
			json.Unmarshal([]byte(resBody), &response)

			assert.Equal(t, testCase.expectCode, res.Code)
			assert.Equal(t, testCase.responseStatus, response.Status)
			assert.Equal(t, testCase.responseMessage, response.Message)
		}
	}
}