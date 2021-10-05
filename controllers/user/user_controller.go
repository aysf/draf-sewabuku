package user

import (
	"net/http"
	"sewabuku/database"
	"sewabuku/middlewares"
	"sewabuku/models"
	"sewabuku/util"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	userModel database.UserModel
}

// NewController is function to initialize new controller
func NewController(userModel database.UserModel) *Controller {
	return &Controller{
		userModel,
	}
}

// RegisterUserController is controller for user registration
func (controller *Controller) RegisterUserController(c echo.Context) error {
	var userRequest models.User
	c.Bind(&userRequest)

	if err := c.Validate(&userRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, util.ResponseError("Check Your Input", nil))
	}

	user := models.User{
		Name:             userRequest.Name,
		OrganizationName: userRequest.OrganizationName,
		Email:            userRequest.Email,
		Password:         userRequest.Password,
		Address:          userRequest.Address,
	}

	_, err := controller.userModel.Register(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Register Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Register Success", nil))
}

// LoginUserController is controller for user login
func (controller *Controller) LoginUserController(c echo.Context) error {
	var userRequest models.User
	c.Bind(&userRequest)

	if err := c.Validate(&userRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, util.ResponseError("Check Your Input", nil))
	}

	user, err := controller.userModel.Login(userRequest.Email, userRequest.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Login Failed", nil))
	}

	token := struct {
		Token string `json:"token"`
	}{
		Token: user.Token,
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Login Success", token))
}

// GetUserProfileController is controller for get user profile
func (controller *Controller) GetUserProfileController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := controller.userModel.GetProfile(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get User Profile", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get User Profile", user))
}

// UpdateUserProfileController is controller for user edit their profile
func (controller *Controller) UpdateUserProfileController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)
	var userRequest models.User
	c.Bind(&userRequest)

	if err := c.Validate(&userRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, util.ResponseError("Check Your Input", nil))
	}

	user := models.User{
		Name:    userRequest.Name,
		Email:   userRequest.Email,
		Address: userRequest.Address,
	}

	_, err := controller.userModel.UpdateProfile(user, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Update User Profile", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Update Profile", nil))
}

// UpdatePasswordController is controller for user edit their password
func (controller *Controller) UpdatePasswordController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	var userRequest models.User
	c.Bind(&userRequest)

	if err := c.Validate(&userRequest); err != nil {
		return c.JSON(http.StatusInternalServerError, util.ResponseError("Check Your Input", nil))
	}

	user := models.User{Password: userRequest.Password}

	if _, err := controller.userModel.UpdatePassword(user, userId); err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Change Password", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Change Password", nil))
}
