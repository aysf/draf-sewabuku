package book

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sewabuku/database"
	"strconv"
)

type Controller struct {
	bookModel database.BookModel
}

func NewController(bookModel database.BookModel) *Controller {
	return &Controller{
		bookModel,
	}
}

func (controller *Controller) GetAllBookController(c echo.Context) error {
	book, err := controller.bookModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	return c.JSON(http.StatusOK, book)
}

func (controller *Controller) GetBookController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	book, err := controller.bookModel.Get(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	return c.JSON(http.StatusOK, book)
}