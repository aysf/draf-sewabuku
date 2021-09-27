package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sewabuku/database"
	"strconv"
)

type Controller struct {
	userModel database.UserModel
}

func NewController(userModel database.UserModel) *Controller {
	return &Controller{
		userModel,
	}
}

func (controller *Controller) GetAllUserController(c echo.Context) error {
	user, err := controller.userModel.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *Controller) GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	user, err := controller.userModel.Get(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	return c.JSON(http.StatusOK, user)
}