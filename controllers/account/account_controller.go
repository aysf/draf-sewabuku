package account

import (
	"fmt"
	"net/http"
	"sewabuku/database"
	"sewabuku/middlewares"
	"sewabuku/models"
	"sewabuku/util"
	"strconv"

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

func (controller *Controller) BalanceTransaction(c echo.Context) error {
	code, _ := strconv.Atoi(c.QueryParam("code"))

	var entryRequest models.Entry
	c.Bind(&entryRequest)

	// if err := c.Bind(&entryRequest); err != nil {
	// 	fmt.Println("error is", err)
	// 	return c.JSON(http.StatusBadRequest, "fail")
	// }

	userId := middlewares.ExtractTokenUserId(c)
	fmt.Println(entryRequest)
	var amount int
	if code == 1 {
		amount = entryRequest.Amount
	} else if code == 2 {
		amount = -1 * entryRequest.Amount
	}

	entry := models.Entry{
		AccountID: uint(userId),
		Amount:    amount,
	}

	_, err := controller.accountModel.Transaction(entry)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	var message string
	if code == 1 {
		message = "Deposit success"
	} else if code == 2 {
		message = "Withdrawal success"
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess(message, entry))

}
