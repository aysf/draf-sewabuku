package main

import (
	"sewabuku/config"
	"sewabuku/controllers/account"
	"sewabuku/controllers/book"
	"sewabuku/controllers/cart"
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
	bookModel := database.NewBookModel(db)
	userModel := database.NewUserModel(db)
	accountModel := database.NewAccountModel(db)
	cartModel := database.NewCartModel(db)

	config.InsertDumyData(db)

	// Initialize controller
	newUserController := user.NewController(userModel)
	newAccountController := account.NewController(accountModel)
	newBookController := book.NewBookController(bookModel)
	newCartController := cart.NewCartController(cartModel)

	// API path and controller
	routes.UserPath(e, newUserController)
	routes.BookPath(e, newBookController)
	routes.AccountPath(e, newAccountController)
	routes.CartPath(e, newCartController)

	// run server
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
