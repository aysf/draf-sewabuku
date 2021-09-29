package routes

import (
	"sewabuku/controllers/account"

	"github.com/labstack/echo/v4"
)

func AccountPath(e *echo.Echo, accountController *account.Controller) {
	accountGroup := e.Group("/account")

	accountGroup.GET("", accountController.ShowAccountBalance)

	accountGroup.GET("", accountController.AddEntries)

}
