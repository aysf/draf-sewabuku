package routes

import (
	"sewabuku/controllers"

	"github.com/labstack/echo/v4"
)

func Setup(e *echo.Echo) {

	// authentication
	e.POST("/api/register", controllers.Register)
	e.POST("/api/login", controllers.Login)
	e.GET("/api/user", controllers.User)
	e.POST("/api/logout", controllers.Logout)

	// book service
	e.GET("/", controllers.GetUsers)
}
