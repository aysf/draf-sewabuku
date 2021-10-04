package book

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sewabuku/database"
	"sewabuku/middlewares"
	"sewabuku/models"
	"sewabuku/util"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
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

const ExtensionAllowed = ".jpg, .jpeg, .png"

// func (controller *Controller) GetAllBookController(c echo.Context) error {
// 	book, err := controller.bookModel.GetAll()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "fail")
// 	}

// 	return c.JSON(http.StatusOK, book)
// }

func (h *Controller) GetAllBooks(c echo.Context) error {
	books, err := h.bookModel.GetAllBooks()
	if err != nil {
		Response := util.ResponseError("failed error", nil)
		return c.JSON(http.StatusBadRequest, Response)
	}

	responseBook := FormatResponseBooks(books)

	response := util.ResponseSuccess("success get all books", responseBook)
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

	responseFormat := FormatDetailsBook(books)

	response := util.ResponseSuccess("success", responseFormat)
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
	name := c.Param("name")

	books, err := h.bookModel.GetByNameBook(name)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if len(books) == 0 {
		response := util.ResponseFail(fmt.Sprintf("there's no book with name %v", name), nil)
		return c.JSON(http.StatusOK, response)
	}

	responseBook := FormatResponseBooks(books)

	response := util.ResponseSuccess("ok", responseBook)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) GetByCategoryID(c echo.Context) error {
	category_id, err := strconv.Atoi(c.Param("id"))
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

	responseBook := FormatResponseBooks(books)

	response := util.ResponseSuccess("success", responseBook)
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
	author_id, err := strconv.Atoi(c.Param("id"))
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

	responseBook := FormatResponseBooks(books)

	response := util.ResponseSuccess("success", responseBook)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) GetByPublisherID(c echo.Context) error {
	publisher_id, err := strconv.Atoi(c.Param("id"))
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
	responseBook := FormatResponseBooks(books)

	response := util.ResponseSuccess("success", responseBook)
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
	check, err := h.bookModel.CheckPublisherName(name)
	if err != nil {
		response := util.ResponseError("failed error", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if !check {
		response := util.ResponseFail("cannot input name author with same name which already exist", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	var input models.Publisher
	input.Name = name

	publisher, err := h.bookModel.CreateNewPublisher(input)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := util.ResponseSuccess("successfully create new publisher", publisher)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) CreateNewAuthor(c echo.Context) error {

	_ = middlewares.ExtractTokenUserId(c)
	name := c.FormValue("name")
	name = strings.ToLower(name)
	name = strings.TrimSpace(fmt.Sprintf("%v", name))

	aneh := strings.ContainsAny(name, "}{!@#$%^&*)''?(-=_/\\+`~][|.,;:")
	if aneh {
		response := util.ResponseError("cannot create new author if there are special characters", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	check, err := h.bookModel.CheckAuthorName(name)
	if err != nil {
		response := util.ResponseError("failed error", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if !check {
		response := util.ResponseFail("cannot input name author with same name which already exist", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	var input models.Author
	input.Name = name
	author, err := h.bookModel.CreateNewAuthor(input)
	if err != nil {
		response := util.ResponseError("error", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	response := util.ResponseSuccess("successfully create new publisher", author)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) BorrowBook(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)

	bookID, _ := strconv.Atoi(c.Param("id"))

	check, err := h.bookModel.GetBookByID(uint(bookID))
	if err != nil {
		response := util.ResponseError("failed", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	if check.Quantity == 0 {
		response := util.ResponseFail("sorry someone is borrowing this book, please wait until it gets returned ", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	carts, err := h.bookModel.BorrowBook(bookID, user_id)
	if err != nil {
		response := util.ResponseError("failed to borrow book", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	response := util.ResponseSuccess("successfully asking for borrow books", carts)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) InsertBook(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)

	ya, err := h.bookModel.CheckBorrowBook(user_id)
	if !ya || err != nil {
		response := util.ResponseError("can not insert new book if you are still borrowing someone`s book", nil)
		return c.JSON(http.StatusInternalServerError, response)
	}
	var input models.InputBook

	err = c.Bind(&input)
	if err != nil {
		response := util.ResponseError("error internal", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	if input.Tittle == "" || input.PublishYear == 0 {
		response := util.ResponseFail("please input name of your book and year of publishment of your book", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if input.CategoryID == 0 {
		input.CategoryID = 1
	}
	if input.AuthorID == 0 {
		input.AuthorID = 1
	}
	if input.PublisherID == 0 {
		input.PublisherID = 1
	}

	books := models.BookData{
		Tittle:      input.Tittle,
		CategoryID:  input.CategoryID,
		UserID:      uint(user_id),
		AuthorID:    input.AuthorID,
		PublisherID: input.PublisherID,
		PublishYear: input.PublishYear,
		Price:       uint(input.Price),
		Quantity:    input.Quantity,
		Description: input.Description,
	}

	book, err := h.bookModel.InsertNewBook(books)
	if err != nil {
		response := util.ResponseError("failed to insert book", err)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	formatResponse := FormatDetailsBook(book)
	response := util.ResponseSuccess("successfully insert your new book", formatResponse)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) UpdatePhotoBook(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)
	bookID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return fmt.Errorf("error internal guys")
	}

	book, err := h.bookModel.GetBookByID(uint(bookID))
	if err != nil {
		response := util.ResponseError("failed to update photo of the book", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	if book.UserID != uint(user_id) {
		response := util.ResponseFail("you are not owner of this book", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	err = c.Request().ParseMultipartForm(1024)
	if err != nil {
		response := util.ResponseError("failed to update photo of the book", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	foto, file, err := c.Request().FormFile("file")
	if err != nil {
		resp := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusInternalServerError, resp)

	}

	defer foto.Close()

	filebyte, err := ioutil.ReadAll(foto)
	if err != nil {
		resp := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	mime := mimetype.Detect(filebyte)
	if strings.Index(ExtensionAllowed, mime.Extension()) == -1 {
		response := util.ResponseError("file type extension is not allowed", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	directory, err := os.Getwd()
	if err != nil {
		resp := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	fotofileName := fmt.Sprintf("/%d,%s,%s", bookID, file.Filename, mime.Extension())
	book, err = h.bookModel.UpdatePhoto(fotofileName, bookID)
	if err != nil {
		resp := util.ResponseError("internal error", err)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	pathLocation := filepath.Join(directory, "image", fotofileName)
	targetFile, err := os.OpenFile(pathLocation, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		resp := util.ResponseError("internal error", err)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	defer targetFile.Close()

	response := util.ResponseSuccess("successfully update book photo", book)
	return c.JSON(http.StatusOK, response)

}
