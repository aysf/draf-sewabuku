package main

import (
	"sewabuku/config"
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

	// Initialize controller
	newUserController := user.NewController(userModel)
	newBookController := book.NewController(bookModel)

	// API path and controller
	routes.UserPath(e, newUserController)
	routes.BookPath(e, newBookController)

	// run server
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
