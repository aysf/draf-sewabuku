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
			reqBody:         map[string]string{"name": "", "email": "agito@rider.jp", "password": "kuuga99"},
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

	e := echo.New()

	modelUser := database.NewUserModel(db)

	controllerUser := NewController(modelUser)

	for _, testCase := range testCases {
		body, _ := json.Marshal(testCase.reqBody)
		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
		fmt.Println("--------------------------------------")
		fmt.Println(req.Body)
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