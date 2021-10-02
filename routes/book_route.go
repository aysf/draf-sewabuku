package routes

import (
	"os"
	"sewabuku/controllers/book"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BookPath(e *echo.Echo, bookController *book.Controller) {
	bookGroup := e.Group("/books")
	jwtAuth := e.Group("/books")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))
	// bookGroup.GET("", bookController.GetAllBookController)

	bookGroup.GET("/s/:keyword", bookController.SearchBookController)

	bookGroup.GET("/", bookController.GetAllBooks)

	bookGroup.GET("/category/:id", bookController.GetByCategoryID)

	//bookGroup.DELETE("/:id", bookController.DeleteBookController)
}
