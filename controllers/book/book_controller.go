package book

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

type ControllerBook struct {
	service database.RepositoryBook
}

func NewBookController(service database.RepositoryBook) *ControllerBook {
	return &ControllerBook{service: service}
}

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

func (h *ControllerBook) InsertBook(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	var input models.InputBook

	err := c.Bind(&input)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	if input.Price == 0 {
		response := util.ResponseFail("cannot insert book if not fill up price column", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	var bookdata models.BookData
	bookdata.OwnerID = uint(id)
	bookdata.CategoryID = input.CategoryID
	bookdata.PublishDate = input.PublishDate
	bookdata.Title = input.Title
	bookdata.Author = input.Author
	bookdata.Publisher = input.Publisher
	bookdata.PeiceBook = input.Price

	err = h.service.InputBook(bookdata)
	if err != nil {
		response := util.ResponseFail(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := util.ResponseSuccess("successfully insert book", bookdata)
	return c.JSON(http.StatusOK, response)

}

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

func (h *ControllerBook) UpdateBook(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)

	bookid, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		response := util.ResponseError("error internal", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	book, err := h.service.GetBookByID(uint(bookid))
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if book.OwnerID != uint(id) {
		response := util.ResponseFail("not owner of this book", nil)
		return c.JSON(http.StatusUnauthorized, response)
	}

	var input models.InputBook

	err = c.Bind(&input)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	if input.Price == 0 {
		response := util.ResponseFail("cannot insert book if not fill up price column", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	book.CategoryID = input.CategoryID
	book.PublishDate = input.PublishDate
	book.Title = input.Title
	book.Author = input.Author
	book.Publisher = input.Publisher
	book.PeiceBook = input.Price

	err = h.service.UpdateBook(book)
	if err != nil {
		response := util.ResponseFail(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := util.ResponseSuccess("successfully insert book", book)
	return c.JSON(http.StatusOK, response)

}

// func (h *ControllerBook) BorrowBook(c echo.Context) error {
// 	id := middlewares.ExtractTokenUserId(c)

// 	bookid, err := strconv.Atoi(c.QueryParam("id"))
// 	if err != nil {
// 		response := util.ResponseError("error internal", nil)
// 		return c.JSON(http.StatusInternalServerError, response)
// 	}
// 	id := c.QueryParam("id")

// 	book, err := h.service.GetBookByID(id)
// 	if err != nil {
// 		response := util.ResponseError(err.Error(), nil)
// 		return c.JSON(http.StatusUnprocessableEntity, response)
// 	}

// }
