package book

// import (
// 	"net/http"
// 	"sewabuku/config"
// 	"sewabuku/controllers/book"
// 	"testing"
// )

// // import "testing"

// func TestingGetByKeyWord(t *testing.T) {
// 	testCases := []struct {
// 		testName   string
// 		category   int
// 		author     int
// 		publisher  int
// 		expectCode int
// 		expectMsg  string
// 	}{
// 		{
// 			testName:   "success",
// 			category:   2,
// 			publisher:  1,
// 			author:     3,
// 			expectCode: http.StatusOK,
// 			expectMsg:  "success",
// 		}, {
// 			testName:   "success",
// 			category:   3,
// 			author:     2,
// 			publisher:  2,
// 			expectCode: http.StatusOK,
// 			expectMsg:  "success",
// 		},
// 	}

// 	db := config.DBConnect()
// 	bookHandler := book.NewBookController(db)

// }
