package routes

import (
	"github.com/labstack/echo/v4"
	"sewabuku/controllers/user"
)

func UserPath(e *echo.Echo, userController *user.Controller) {
	// ------------------------------------------------------------------
	// Login & register
	// ------------------------------------------------------------------
	//e.POST("/users/register", userController.RegisterUserController)
	//e.POST("/users/login", userController.LoginUserController)

	// ------------------------------------------------------------------
	// CRUD User
	// ------------------------------------------------------------------
	e.GET("/users", userController.GetAllUserController)
	e.GET("/users/:id", userController.GetUserController)
	//e.PUT("/users/:id", userController.EditUserController)
	//e.DELETE("/users/:id", userController.DeleteUserController)
}