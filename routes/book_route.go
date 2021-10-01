package routes

import (
	"sewabuku/controllers/book"

	"github.com/labstack/echo/v4"
)

func BookPath(e *echo.Echo, bookController *book.Controller) {
	bookGroup := e.Group("/books")

	bookGroup.GET("", bookController.GetAllBookController)

	bookGroup.GET("/:id", bookController.GetBookController)

	bookGroup.GET("/s/:keyword", bookController.SearchBookController)

	//bookGroup.PUT("/:id", bookController.EditBookController)

	//bookGroup.DELETE("/:id", bookController.DeleteBookController)
}
