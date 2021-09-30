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

	bookGroup.GET("/category", bookController.GetByCategory)

	bookGroup.GET("/name", bookController.GetBookByname)

	bookGroup.GET("/author", bookController.GetByAuthor)

	bookGroup.GET("/publisher", bookController.GetByPublisher)

	jwtAuth.POST("/book", bookController.InsertBook)

	jwtAuth.PUT("/book", bookController.UpdateBook)

	jwtAuth.DELETE("/book", bookController.DeleteBook)

}
