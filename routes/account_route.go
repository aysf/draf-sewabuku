package routes

import (
	"sewabuku/controllers/book"

	"github.com/labstack/echo/v4"
)

func AccountPath(e *echo.Echo, bookController *book.Controller) {
	bookGroup := e.Group("/books")

	bookGroup.GET("", bookController.GetAllBookController)

	bookGroup.GET("/:id", bookController.GetBookController)

	//bookGroup.PUT("/:id", bookController.EditBookController)

	//bookGroup.DELETE("/:id", bookController.DeleteBookController)
}
