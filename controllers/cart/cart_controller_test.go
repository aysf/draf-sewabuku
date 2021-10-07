package cart

import (
	"fmt"
	"net/http"
	"sewabuku/config"
	"sewabuku/models"
	"testing"
)

func TestController_RentBook(t *testing.T) {
	var testCases = []struct {
		name            string
		reqBody         string
		expectCode      int
		responseStatus  string
		responseMessage string
	}{
		{
			name:            "case 1: rent available books",
			reqBody:         `{"firstName": "Ananto", "lastName":"Wicaksono", "email": "aw@test.com", "password": "ay123"}`,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 2: rent books when does not have enough balance",
			reqBody:         ``,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 3: rent the book id that already rented before",
			reqBody:         ``,
			expectCode:      http.StatusOK,
			responseStatus:  "success",
			responseMessage: "Register Success",
		},
		{
			name:            "case 4: rent book for user itself",
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
		&models.BookData{},
	)
	db.AutoMigrate(
		&models.Account{},
		&models.User{},
	)

	fmt.Print(db, testCases)
}
