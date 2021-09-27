package routes

import (
	"github.com/labstack/echo/v4"
	"sewabuku/controllers/user"
)

func UserPath(e *echo.Echo, userController *user.Controller) {
	groupUser := e.Group("/users")
	groupUser.POST("/users/register", userController.RegisterUserController)
	groupUser.POST("/users/login", userController.LoginUserController)
	groupUser.GET("", userController.GetUserProfileController)
}