package routes

import (
	"os"
	"sewabuku/controllers/account"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AccountPath(e *echo.Echo, accountController *account.Controller) {
	jwtAuth := e.Group("/account")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))

	jwtAuth.POST("/transaction", accountController.TopupWithdraw)
	jwtAuth.PUT("/deposit", accountController.DepositTransfer)
}
