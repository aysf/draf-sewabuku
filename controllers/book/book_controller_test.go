package book

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"sewabuku/config"
// 	"sewabuku/database"
// 	"sewabuku/models"
// 	"testing"

// 	"github.com/labstack/echo/v4"
// 	"github.com/tj/assert"
// )

// // import "testing"

// func TestingGetByCategory(t *testing.T) {
// 	var testCases = []struct {
// 		name           string
// 		reqBody        string
// 		expectCode     int
// 		responseStatus string
// 	}{
// 		{
// 			name:           "success",
// 			reqBody:        "pengetahuan",
// 			expectCode:     200,
// 			responseStatus: "success",
// 		},
// 		{
// 			name:           "failed",
// 			reqBody:        "alam",
// 			expectCode:     200,
// 			responseStatus: "failed",
// 		},
// 	}
// 	db := config.DBConnectTest()

// 	db.AutoMigrate(&models.User{})

// 	e := echo.New()

// 	for _, tescase := range testCases {
// 		fmt.Println("======================================")
// 		fmt.Println("ini test ke", i)
// 		// Create new repo and controller user
// 		modelUser := database.NewUserModel(db)
// 		controllerUser := NewController(modelUser)
// 		body, _ := json.Marshal(testCase.reqBody)
// 		req := httptest.NewRequest(http.MethodPost, "/users/register", bytes.NewBuffer(body))
// 		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 		res := httptest.NewRecorder()
// 		ctx := e.NewContext(req, res)

// 		if testCase.name == "test3" {
// 			db.Migrator().DropTable(&models.User{})
// 		}

// 		if assert.NoError(t, controllerUser.RegisterUserController(ctx)) {
// 			resBody := res.Body.String()

// 			var response userResponse
// 			json.Unmarshal([]byte(resBody), &response)

// 			assert.Equal(t, testCase.expectCode, res.Code)
// 			assert.Equal(t, testCase.responseStatus, response.Status)
// 			assert.Equal(t, testCase.responseMessage, response.Message)
// 		}
// 	}
// }
