package routes

import (
	"os"
	"sewabuku/controllers/user"

	"github.com/labstack/echo/v4/middleware"

	"github.com/labstack/echo/v4"
)

func UserPath(e *echo.Echo, userController *user.Controller) {
	noAuth := e.Group("/users")
	jwtAuth := e.Group("/users")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))

	noAuth.POST("/register", userController.RegisterUserController)

	noAuth.POST("/login", userController.LoginUserController)

	jwtAuth.GET("/profile", userController.GetUserProfileController)

	jwtAuth.PUT("/profile", userController.UpdateUserProfileController)

	jwtAuth.PUT("/change-password", userController.UpdatePasswordController)

	jwtAuth.PUT("/logout", userController.LogoutUserController)

	jwtAuth.GET("/borrowed", userController.GetBorrowedController)

	jwtAuth.GET("/lent", userController.GetLentController)

	jwtAuth.PUT("/book-rating/:id/", userController.InsertRatingBookController)

	jwtAuth.PUT("/borrower-rating/:id/", userController.InsertRatingBorrowerController)
}
