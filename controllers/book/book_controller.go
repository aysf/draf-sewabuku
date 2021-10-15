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

func (h *Controller) SearchAll(c echo.Context) error {
	keyword := "%" + c.Param("keyword") + "%"
	publisher, _ := strconv.Atoi(c.QueryParam("publisher"))
	category, _ := strconv.Atoi(c.QueryParam("category"))
	author, _ := strconv.Atoi(c.QueryParam("author"))

	books, err := h.bookModel.SearchBooks(keyword, author, publisher, category)
	if err != nil {
		Response := util.ResponseError("failed to get book", err)
		return c.JSON(http.StatusBadRequest, Response)
	}

	if len(books) == 0 {
		resp := util.ResponseFail("theres no book found", nil)
		return c.JSON(http.StatusOK, resp)
	}
	formatResponse := FormatResponseBooks(books)
	response := util.ResponseSuccess("success get books", formatResponse)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) GetDetailsBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		Response := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusBadRequest, Response)
	}

	books, _ := h.bookModel.GetBookByID(uint(id))
	// if err != nil {
	// 	Response := util.ResponseError("cann not get book", nil)
	// 	return c.JSON(http.StatusBadRequest, Response)
	// }

	if books.ID == 0 {
		Response := util.ResponseFail("there's no book found", nil)
		return c.JSON(http.StatusBadRequest, Response)
	}

	responseFormat := FormatDetailsBook(books)

	response := util.ResponseSuccess("success get details book", responseFormat)
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

func (h *Controller) FilterAuthorCategoryPublisher(c echo.Context) error {
	publisher, _ := strconv.Atoi(c.QueryParam("publisher"))
	category, _ := strconv.Atoi(c.QueryParam("category"))
	author, _ := strconv.Atoi(c.QueryParam("author"))

	books, err := h.bookModel.GetByKeywordID(author, category, publisher)
	if err != nil {
		response := util.ResponseError(err.Error(), nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	if len(books) == 0 {
		response := util.ResponseFail("no book found", nil)
		return c.JSON(http.StatusOK, response)

	}
	responseBook := FormatResponseBooks(books)

	response := util.ResponseSuccess("success get books", responseBook)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) CreateNewPublisher(c echo.Context) error {
	_ = middlewares.ExtractTokenUserId(c)
	name := c.FormValue("name")
	if name == "" {
		response := util.ResponseError("please input publisher name", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	name = strings.ToLower(name)
	name = strings.TrimSpace(fmt.Sprintf("%v", name))

	aneh := strings.ContainsAny(name, "}{!@#$%^&*)''?(-=_/\\+`~][|.,;:")
	if aneh {
		response := util.ResponseError("cannot create new author if there are special characters", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}
	check, _ := h.bookModel.CheckPublisherName(name)
	// if err != nil {
	// 	response := util.ResponseError("failed error", nil)
	// 	return c.JSON(http.StatusUnprocessableEntity, response)
	// }

	if !check {
		response := util.ResponseFail("cannot input name author with same name which already exist", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	var input models.Publisher
	input.Name = name

	publisher, _ := h.bookModel.CreateNewPublisher(input)
	// if err != nil {
	// 	response := util.ResponseError(err.Error(), nil)
	// 	return c.JSON(http.StatusUnprocessableEntity, response)
	// }

	response := util.ResponseSuccess("successfully create new publisher", publisher)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) CreateNewAuthor(c echo.Context) error {

	_ = middlewares.ExtractTokenUserId(c)
	name := c.FormValue("name")
	if name == "" {
		response := util.ResponseError("please input author name", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	name = strings.ToLower(name)
	name = strings.TrimSpace(fmt.Sprintf("%v", name))

	aneh := strings.ContainsAny(name, "}{!@#$%^&*)''?(-=_/\\+`~][|.,;:")
	if aneh {
		response := util.ResponseError("cannot create new author if there are special characters", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)

	}

	check, _ := h.bookModel.CheckAuthorName(name)
	// if err != nil {
	// 	response := util.ResponseError("failed error", nil)
	// 	return c.JSON(http.StatusUnprocessableEntity, response)
	// }

	if !check {
		response := util.ResponseFail("cannot input name author with same name which already exist", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	var input models.Author
	input.Name = name
	author, _ := h.bookModel.CreateNewAuthor(input)
	// if err != nil {
	// 	response := util.ResponseError("error", nil)
	// 	return c.JSON(http.StatusUnprocessableEntity, response)

	// }

	response := util.ResponseSuccess("successfully create new author", author)
	return c.JSON(http.StatusOK, response)
}

func (h *Controller) InsertBook(c echo.Context) error {
	user_id := middlewares.ExtractTokenUserId(c)

	ya, err := h.bookModel.CheckBorrowBook(user_id)
	if !ya || err != nil {
		response := util.ResponseError("can not insert new book if you are still borrowing someone`s book", nil)
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	var input models.InputBook

	err = c.Bind(&input)
	if err != nil {
		response := util.ResponseError("error internal", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	if input.Title == "" || input.PublishYear == 0 {
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
		Title:       input.Title,
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
	bookID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := util.ResponseError("error internal", nil)
		return c.JSON(http.StatusInternalServerError, response)
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

	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("error nangkep file")
		response := util.ResponseFail("failed to update books's photo", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	src, err := file.Open()
	if err != nil {
		response := util.ResponseFail("internal error", err.Error())
		return c.JSON(http.StatusUnprocessableEntity, response)
	}
	defer src.Close()

	filebyte, err := ioutil.ReadAll(src)
	if err != nil {
		resp := util.ResponseError("internal error", err.Error())
		return c.JSON(http.StatusInternalServerError, resp)
	}

	mime := mimetype.Detect(filebyte)
	if strings.Index(ExtensionAllowed, mime.Extension()) == -1 {
		response := util.ResponseError("file type extension is not allowed", fmt.Sprintf("your extension %s", mime.Extension()))
		return c.JSON(http.StatusUnprocessableEntity, response)
	}

	directory, err := os.Getwd()
	if err != nil {
		resp := util.ResponseError("internal error", nil)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	fotofileName := fmt.Sprintf("/%d,%s", bookID, file.Filename)
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

	getnewbook, err := h.bookModel.GetBookByID(uint(bookID))
	if err != nil {
		resp := util.ResponseError("error get book but book's photo has been updated", err)
		return c.JSON(http.StatusInternalServerError, resp)
	}

	formatresponse := FormatDetailsBook(getnewbook)

	response := util.ResponseSuccess("successfully update book photo", formatresponse)
	return c.JSON(http.StatusOK, response)

}

func (h *Controller) GetCommentBookID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response := util.ResponseError("error internal", err)
		return c.JSON(http.StatusInternalServerError, response)
	}

	commentBooks, err := h.bookModel.GetCommentBookID(id)
	if err != nil {
		rsponse := util.ResponseError("cant get comment", err)
		return c.JSON(http.StatusUnprocessableEntity, rsponse)
	}

	if len(commentBooks) == 0 {
		rsponse := util.ResponseSuccess("theres no any comment for this book yet", nil)
		return c.JSON(http.StatusOK, rsponse)
	}

	response := util.ResponseSuccess("successfully get book comments", commentBooks)
	return c.JSON(http.StatusOK, response)
}
