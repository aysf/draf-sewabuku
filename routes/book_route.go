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

	bookGroup.GET("/all", bookController.GetAllBooks)

	bookGroup.GET("/category", bookController.GetByCategoryID)

	bookGroup.GET("/author", bookController.GetByAuthorID)

	bookGroup.GET("/publisher", bookController.GetByPublisherID)

	bookGroup.GET("/listauthor", bookController.GetListAuthor)

	bookGroup.GET("/listcategory", bookController.GetListCategory)

	bookGroup.GET("/listpublisher", bookController.GetListPublisher)

	bookGroup.GET("/name", bookController.GetBookByname)

	bookGroup.GET("/details", bookController.GetDetailsBook)

	jwtAuth.POST("/newauthor", bookController.CreateNewAuthor)

	jwtAuth.POST("/newpublisher", bookController.CreateNewPublisher)

	jwtAuth.POST("/borrow", bookController.BorrowBook)
}
