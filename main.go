package main

import (
	"sewabuku/config"
	"sewabuku/controllers/account"
	"sewabuku/controllers/book"
	"sewabuku/controllers/user"
	"sewabuku/database"
	m "sewabuku/middlewares"
	"sewabuku/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize database connection
	db := config.DBConnect()

	// Create echo http
	e := echo.New()

	// Initialize model
	userModel := database.NewUserModel(db)
	bookModel := database.NewBookModel(db)
	accountModel := database.NewAccountModel(db)
	config.InsertDumyData(db)

	// Initialize controller
	newUserController := user.NewController(userModel)
	newBookController := book.NewController(bookModel)
	newAccountController := account.NewController(accountModel)

	// API path and controller
	routes.UserPath(e, newUserController)
	routes.BookPath(e, newBookController)
	routes.AccountPath(e, newAccountController)

	// run server
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
