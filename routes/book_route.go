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

	bookGroup.GET("/search/:keyword", bookController.SearchAll)

	bookGroup.GET("/all", bookController.GetAllBooks)

	bookGroup.GET("/s", bookController.FilterAuthorCategoryPublisher)

	bookGroup.GET("/listauthor", bookController.GetListAuthor)

	bookGroup.GET("/listcategory", bookController.GetListCategory)

	bookGroup.GET("/listpublisher", bookController.GetListPublisher)

	bookGroup.GET("/details/:id", bookController.GetDetailsBook)

	jwtAuth.POST("/newauthor", bookController.CreateNewAuthor)

	jwtAuth.POST("/newpublisher", bookController.CreateNewPublisher)

	jwtAuth.PUT("/bookphoto/:id", bookController.UpdatePhotoBook)

	jwtAuth.POST("/newbook", bookController.InsertBook)
}
