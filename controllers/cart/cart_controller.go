package cart

import (
	"net/http"
	"sewabuku/database"
	"sewabuku/middlewares"
	"sewabuku/models"
	"sewabuku/util"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	cartModel database.CartModel
}

// NewController is function to initialize new controller
func NewCartController(cartModel database.CartModel) *Controller {
	return &Controller{
		cartModel,
	}
}

// RegisterUserController is controller for user registration
func (controller *Controller) RentBook(c echo.Context) error {
	var cartRequest models.Cart
	c.Bind(&cartRequest)

	userId := middlewares.ExtractTokenUserId(c)
	cart := models.Cart{
		UserID:     uint(userId),
		BookUserID: uint(cartRequest.BookUserID),
		DateLoan:   cartRequest.DateLoan,
		DateDue:    cartRequest.DateDue,
	}

	_, err := controller.cartModel.Rent(cart)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Rent Book Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Rent Book Success", nil))
}

func (controller *Controller) ReturnBook(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	var Date models.Cart
	c.Bind(&Date)

	updateDate := models.Cart{
		DateReturn: Date.DateReturn,
	}

	returnBook, err := controller.cartModel.Return(updateDate.DateReturn, userId, int(Date.BookUserID))

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Return Book", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Return Book", returnBook))
}

func (controller *Controller) ListBook(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	carts, err := controller.cartModel.List(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get List", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success get book borrowing list", carts))

}
