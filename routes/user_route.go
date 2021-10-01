package routes

import (
	"os"
	"sewabuku/controllers/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserPath(e *echo.Echo, userController *user.Controller) {
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))

	e.POST("/users/register", userController.RegisterUserController)

	e.POST("/users/login", userController.LoginUserController)

	jwtAuth.GET("/users/profile", userController.GetUserProfileController)

	jwtAuth.POST("/users/profile", userController.UpdateUserProfileController)

	jwtAuth.POST("/users/change-password", userController.UpdatePasswordController)
}
