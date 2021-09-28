package routes

import (
	"github.com/labstack/echo/v4"
	"sewabuku/controllers/user"
)

func UserPath(e *echo.Echo, userController *user.Controller) {
	e.POST("/users/register", userController.RegisterUserController)

	e.POST("/users/login", userController.LoginUserController)

	e.GET("/users", userController.GetUserProfileController)
}