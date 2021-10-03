package book

import (
	"fmt"
	"net/http"
	"sewabuku/database"
	"sewabuku/middlewares"
	"sewabuku/models"
	"sewabuku/util"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	bookModel database.BookModel
}

func NewBookController(bookModel database.BookModel) *Controller {
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

func (h *Controller) GetAllBooks(c echo.Context) error {
	books, err := h.bookModel.GetAllBooks()
	if err != nil {
		Response := util.ResponseError("failed error", nil)
		return c.JSON(http.StatusBadRequest, Response)
	}

	response := util.ResponseSuccess("success get books", books)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) GetDetailsBook(c echo.Context) error {
	id, err := strconv.Atoi(c.QueryParam("id"))

	if err != nil {
		Response := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusBadRequest, Response)
	}

	books, err := h.bookModel.GetBookByID(uint(id))
	if err != nil {
		Response := util.ResponseError("error disini", nil)
		return c.JSON(http.StatusBadRequest, Response)
	}

	response := util.ResponseSuccess("success", books)
	return c.JSON(http.StatusOK, response)
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

func (h *Controller) GetBookByname(c echo.Context) error {
	name := c.QueryParam("name")

	books, err := h.bookModel.GetByNameBook(name)
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

func (h *Controller) GetByCategoryID(c echo.Context) error {
	category_id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		response := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	books, err := h.bookModel.GetByCategoryID(category_id)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}
	if len(books) == 0 {
		response := util.ResponseFail("there is no book in this category", nil)
		return c.JSON(http.StatusOK, response)
	}

	response := util.ResponseSuccess("success", books)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) GetListAuthor(c echo.Context) error {
	listAuthors, err := h.bookModel.ListAuthor()
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	response := util.ResponseSuccess("success", listAuthors)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) GetListPublisher(c echo.Context) error {
	publishers, err := h.bookModel.GetListPublisher()
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	response := util.ResponseSuccess("success", publishers)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) GetListCategory(c echo.Context) error {
	listCategory, err := h.bookModel.ListCategory()
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	response := util.ResponseSuccess("success", listCategory)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) GetByAuthorID(c echo.Context) error {
	author_id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		response := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	books, err := h.bookModel.GetByAuthorID(author_id)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	if len(books) == 0 {
		response := util.ResponseFail("no book found", nil)
		return c.JSON(http.StatusOK, response)

	}

	response := util.ResponseSuccess("success", books)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) GetByPublisherID(c echo.Context) error {
	publisher_id, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		response := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}

	books, err := h.bookModel.GetByPublisherID(publisher_id)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	if len(books) == 0 {
		response := util.ResponseFail("no book found", nil)
		return c.JSON(http.StatusOK, response)

	}

	response := util.ResponseSuccess("success", books)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) CreateNewPublisher(c echo.Context) error {
	name := c.FormValue("name")
	name = strings.ToLower(name)
	name = strings.TrimSpace(fmt.Sprintf("%v", name))

	aneh := strings.ContainsAny(name, "}{!@#$%^&*)''?(-=_/\\+`~][|.,;:")
	if aneh {
		response := util.ResponseError("cannot create new author if there are special characters", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}
	var input models.Publisher
	input.Name = name

	input, err := h.bookModel.CreateNewPublisher(input)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := util.ResponseSuccess("successfully create new publisher", input)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) CreateNewAuthor(c echo.Context) error {
	name := c.FormValue("name")
	name = strings.ToLower(name)
	name = strings.TrimSpace(fmt.Sprintf("%v", name))

	aneh := strings.ContainsAny(name, "}{!@#$%^&*)''?(-=_/\\+`~][|.,;:")
	if aneh {
		response := util.ResponseError("cannot create new author if there are special characters", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	var input models.Author
	input.Name = name
	author, err := h.bookModel.CreateNewAuthor(input)
	if err != nil {
		response := util.ResponseError("error", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}
	if author.Name == name {
		response := util.ResponseError("cannot create new author with same name which already exist", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	response := util.ResponseSuccess("successfully create new publisher", input)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) UpdatePhotoBook(c echo.Context) error {
	_ = middlewares.ExtractTokenUserId(c)
	_ = c.QueryParam("id")

	// foto, file, err := c.Request().FormFile("file")

	response := util.ResponseError("cannot create new author with same name which already exist", nil)
	return c.JSON(http.StatusUnprocessableEntity, response)

}
