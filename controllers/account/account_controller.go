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

func (controller *Controller) TopupWithdraw(c echo.Context) error {
	code, _ := strconv.Atoi(c.QueryParam("code"))

	var entryRequest models.Entry
	c.Bind(&entryRequest)

	userId := middlewares.ExtractTokenUserId(c)
	accountId := fmt.Sprintf("a-%d", userId)

	var amount int
	if code == 1 {
		amount = entryRequest.Amount
	} else if code == 2 {
		amount = -1 * entryRequest.Amount
	}

	entry := models.Entry{
		AccountID: accountId,
		Amount:    amount,
	}

	_, err := controller.accountModel.Transaction(entry)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	var message string
	if code == 1 {
		message = "Topup success"
	} else if code == 2 {
		message = "Withdrawal success"
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess(message, entry))

}

func (controller *Controller) DepositTransfer(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	var ammountInput models.Account
	c.Bind(&ammountInput)

	ammount := ammountInput.Balance

	ammountUpdate, err := controller.accountModel.UpdateBalance(uint(userId), ammount)
	if err != nil {
		msg := fmt.Sprint(err)
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to make transactoin", msg))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("deposit transfer success", ammountUpdate))
}
