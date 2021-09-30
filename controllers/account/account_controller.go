package account

import (
	"net/http"
	"sewabuku/database"
	"sewabuku/middlewares"
	"sewabuku/models"

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

	return c.JSON(http.StatusOK, account.Balance)
}

func (controller *Controller) AddEntries(c echo.Context) error {
	var entryRequest models.Entry

	if err := c.Bind(&entryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	userId := middlewares.ExtractTokenUserId(c)
	entry := models.Entry{
		AccountID: uint(userId),
		Amount:    entryRequest.Amount,
	}

	_, err := controller.accountModel.Add(entry)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, "success")
}