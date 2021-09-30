package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"sewabuku/models"

	"sewabuku/database"
	"sewabuku/middlewares"
)

type Controller struct {
	userModel database.UserModel
}

func NewController(userModel database.UserModel) *Controller {
	return &Controller{
		userModel,
	}
}

func (controller Controller) RegisterUserController(c echo.Context) error {
	var userRequest models.User

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: userRequest.Password,
	}

	_, err := controller.userModel.Register(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, "success")
}

func (controller Controller) LoginUserController(c echo.Context) error {
	var userRequest models.User

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	user, err := controller.userModel.Login(userRequest.Email, userRequest.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": user.Token,
	})
}

func (controller *Controller) GetUserProfileController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := controller.userModel.GetProfile(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "fail")
	}

	return c.JSON(http.StatusOK, user)
}
