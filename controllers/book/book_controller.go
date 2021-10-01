package book

import (
	"net/http"
	"sewabuku/database"
	"strconv"

	"github.com/labstack/echo/v4"
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

func (controller *Controller) SearchBookController(c echo.Context) error {
	keyword := "%" + c.Param("keyword") + "%"
	category := "%" + c.QueryParam("category") + "%"
	author := "%" + c.QueryParam("author") + "%"

	books, err := controller.bookModel.Search(keyword, author, category)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	return c.JSON(http.StatusOK, books)
}
