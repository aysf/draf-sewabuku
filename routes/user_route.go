package routes

import (
	"github.com/labstack/echo/v4"
	"sewabuku/controllers/user"
)

func UserPath(e *echo.Echo, userController *user.Controller) {
	userGroup := e.Group("/users")

	userGroup.POST("/register", userController.RegisterUserController)

	userGroup.POST("/login", userController.LoginUserController)

	userGroup.GET("", userController.GetUserProfileController)
}