package routes

import (
	"os"
	"sewabuku/controllers/book"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BookPath(e *echo.Echo, bookController *book.ControllerBook) {
	bookGroup := e.Group("/books")
	jwtAuth := e.Group("/books")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))
	// bookGroup.GET("", bookController.GetAllBookController)

	bookGroup.GET("/bycategory", bookController.GetByCategory)

	bookGroup.GET("/namebook", bookController.GetBookByname)

	jwtAuth.POST("/insert", bookController.InsertBook)

	jwtAuth.PUT("/updatebook", bookController.UpdateBook)
}
