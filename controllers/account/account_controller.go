package account

import (
	"net/http"
	"sewabuku/database"
	"sewabuku/middlewares"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	accountModel database.AccountModel
}

func NewController(accountModel database.AccountModel) *Controller {
	return &Controller{
		accountModel,
	}
}

func (controller *Controller) ShowAccountBalance(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	account, err := controller.accountModel.Show(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	return c.JSON(http.StatusOK, account)
}
