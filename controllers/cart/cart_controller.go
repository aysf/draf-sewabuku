package cart

import (
	"fmt"
	"net/http"
	"sewabuku/database"
	"sewabuku/middlewares"
	"sewabuku/models"
	"sewabuku/util"
	"time"

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

	// getting days rental period
	dateLoanInput := cartRequest.DateLoan
	dateDueInput := cartRequest.DateDue
	days := dateDueInput.Sub(dateLoanInput).Hours() / 24

	// getting book and user data
	borrowerBalance, borrowerDeposit, _ := controller.cartModel.GetAccountByUserId(userId)
	book, _ := controller.cartModel.GetBookByBookId(cartRequest.BookDataID)

	rentalFee := int(days) * int(book.Price)

	// check: deposit
	var minDeposit int = int(book.Price) * 90
	if minDeposit > int(borrowerDeposit.Balance) {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Your deposit should be minimum cost at 90 days of rental period", nil))
	}

	// check: balance
	if rentalFee > int(borrowerBalance.Balance) {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Your active balance is not enough, please shorten your rental period or topup your balance", nil))
	}

	// check: user cannot borrow her/his own book
	bookUser, err := controller.cartModel.GetBookByUserId(userId)
	if err != nil {
		return err
	}
	for _, book := range bookUser {
		if book.ID == cartRequest.BookDataID {
			return c.JSON(http.StatusBadRequest, util.ResponseFail("You could not rent your own book", err))
		}
	}

	// check: user cannot borrow the same book id
	bookList, err := controller.cartModel.List(userId)
	if err != nil {
		return err
	}
	var nullTime time.Time
	for _, book := range bookList {
		if book.BookDataID == cartRequest.BookDataID && book.DateReturn == nullTime {
			return c.JSON(http.StatusBadRequest, util.ResponseFail("You could not rent the same book", err))
		}
	}

	// if all check list passed

	// -- update cart
	cart := models.Cart{
		UserID:     uint(userId),
		BookDataID: uint(cartRequest.BookDataID),
		DateLoan:   cartRequest.DateLoan,
		DateDue:    cartRequest.DateDue,
		DateReturn: cartRequest.DateReturn,
	}

	updatedCart, err := controller.cartModel.Rent(cart)
	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Rent Book Failed", err))
	}

	// -- update balance
	rentalFeeCharge := -rentalFee
	if _, err := controller.cartModel.UpdateSaldo(borrowerBalance, rentalFeeCharge); err != nil {
		msg := fmt.Sprintf("error: %s", err)
		return c.JSON(http.StatusInternalServerError, util.ResponseError("Rent Book Failed, balance could not update", msg))
	}

	LenderId, _ := controller.cartModel.GetLenderIdByBookId(cart.BookDataID)
	lenderBalance, _, _ := controller.cartModel.GetAccountByUserId(int(LenderId))
	if _, err := controller.cartModel.UpdateSaldo(lenderBalance, rentalFee); err != nil {
		msg := fmt.Sprintf("error: %s", err)
		return c.JSON(http.StatusInternalServerError, util.ResponseError("Rent Book Failed, balance could not update", msg))
	}
	return c.JSON(http.StatusOK, util.ResponseSuccess("Rent Book Success", updatedCart))
}

func (controller *Controller) ReturnBook(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	var Date models.Cart
	c.Bind(&Date)

	updateDate := models.Cart{
		DateReturn: Date.DateReturn,
	}

	returnBook, err := controller.cartModel.Return(updateDate.DateReturn, userId, int(Date.BookDataID))
	if err != nil {
		msg := fmt.Sprintf("error message: %s", err)
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Return Book", msg))
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

func (controller *Controller) ExtendDateDue(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	var Date models.Cart
	c.Bind(&Date)

	inputDate := models.Cart{
		DateDue: Date.DateDue,
	}

	updateCart, err := controller.cartModel.Extend(inputDate.DateDue, userId, int(Date.BookDataID))

	if err != nil {
		fmt.Println("cek1", err)
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to extend date due", err))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success to extend date due", updateCart))
}
