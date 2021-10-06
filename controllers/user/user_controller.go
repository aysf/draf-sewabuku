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

	if err := c.Validate(userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Check Your Input", nil))
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
		return c.JSON(http.StatusInternalServerError, util.ResponseError("Register Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Register Success", nil))
}

// LoginUserController is controller for user login
func (controller *Controller) LoginUserController(c echo.Context) error {
	var userRequest models.User
	c.Bind(&userRequest)

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

	_, err := controller.userModel.UpdateProfile(userRequest, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Update User Profile", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Update Profile", nil))
}

// UpdatePasswordController is controller for user edit their password
func (controller *Controller) UpdatePasswordController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	var userRequest models.User
	c.Bind(&userRequest)

	user := models.User{Password: userRequest.Password}

	if _, err := controller.userModel.UpdatePassword(user, userId); err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Change Password", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Change Password", nil))
}

// LogoutUserController is controller for user log out
func (controller *Controller) LogoutUserController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	_, err := controller.userModel.Logout(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Logout Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Logout Success", nil))
}

// GetBorrowedController is controller for get borrowed book
func (controller *Controller) GetBorrowedController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	//complete := c.QueryParam("complete")
	complete := c.QueryParam("complete")

	user, err := controller.userModel.GetBorrowed(userId, complete)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get Borrowed Book", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Borrowed Book", user))
}

// GetLentController is controller for get borrowed book
func (controller *Controller) GetLentController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	complete := c.QueryParam("complete")

	user, err := controller.userModel.GetLent(userId, complete)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get Lent Book", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get Lent Book", user))
}