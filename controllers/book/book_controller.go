package book

import (
	"fmt"
	"net/http"
	"sewabuku/database"
	"sewabuku/util"

	"github.com/labstack/echo/v4"
)

type ControllerBook struct {
	service database.RepositoryBook
}

func NewBookController(service database.RepositoryBook) *ControllerBook {
	return &ControllerBook{service: service}
}

// func (controller *ControllerBook) GetAllBookController(c echo.Context) error {
// 	book, err := controller.bookModel.GetAll()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "fail")
// 	}

// 	return c.JSON(http.StatusOK, book)
// }

// func (controller *Controller) GetBookController(c echo.Context) error {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "fail")
// 	}

// 	book, err := controller.bookModel.Get(id)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "fail")
// 	}

// 	return c.JSON(http.StatusOK, book)
// }

func (h *ControllerBook) GetByCategory(c echo.Context) error {
	category := c.QueryParam("category")

	books, err := h.service.GetByCategory(category)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	response := util.ResponseSuccess("success", books)
	return c.JSON(http.StatusOK, response)

}

// func (h *ControllerBook) InsertBook(c echo.Context) error {
// 	var input models.InputBook

// 	err := c.Bind(&input)
// 	if err != nil {
// 		response := util.ResponseError(err.Error(), nil)
// 		return c.JSON(http.StatusUnprocessableEntity, response)

// 	}

// 	id := middlewares.ExtractTokenUserId(c)

// 	var bookdata models.BookData
// 	bookdata.Title = input.Title
// 	bookdata.Author = input.Author
// 	bookdata.Publisher = input.Publisher

// 	var book models.Book
// 	book.Price = input.Price
// 	book.UserID = uint(id)

// 	err = h.service.InputBook(bookdata, book)
// 	if err != nil {
// 		response := util.ResponseError(err.Error(), nil)
// 		return c.JSON(http.StatusUnprocessableEntity, response)

// 	}

// 	response := util.ResponseSuccess("ok", nil)
// 	return c.JSON(http.StatusOK, response)

// }

func (h *ControllerBook) GetBookByname(c echo.Context) error {
	name := c.QueryParam("name")

	books, err := h.service.GetByNameBook(name)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if len(books) == 0 {
		response := util.ResponseSuccess(fmt.Sprintf("there's no book with name %v", name), nil)
		return c.JSON(http.StatusOK, response)
	}

	response := util.ResponseSuccess("ok", books)
	return c.JSON(http.StatusOK, response)

}

// func (h *ControllerBook) BorrowBook(c echo.Context) error {
// 	id := c.QueryParam("id")

// 	book, err := h.service.GetBookByID(id)
// 	if err != nil {
// 		response := util.ResponseError(err.Error(), nil)
// 		return c.JSON(http.StatusUnprocessableEntity, response)
// 	}

// }
