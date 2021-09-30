package routes

import (
	"os"
	"sewabuku/controllers/account"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AccountPath(e *echo.Echo, accountController *account.Controller) {
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))

	jwtAuth.GET("/account", accountController.ShowAccountBalance)

	jwtAuth.POST("/account", accountController.AddEntries)

}
