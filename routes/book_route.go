package routes

import (
	"sewabuku/controllers/book"

	"github.com/labstack/echo/v4"
)

func BookPath(e *echo.Echo, bookController *book.ControllerBook) {
	bookGroup := e.Group("/books")

	// bookGroup.GET("", bookController.GetAllBookController)

	bookGroup.GET("/bycategory", bookController.GetByCategory)

	bookGroup.GET("/namebook", bookController.GetBookByname)

	bookGroup.POST("/insert", bookController.InsertBook)
}
