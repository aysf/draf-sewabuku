package user

import (
	"net/http"
	"os"
	"sewabuku/database"
	"sewabuku/middlewares"
	"sewabuku/models"
	"sewabuku/util"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	bcryptCost, _ := strconv.Atoi(os.Getenv("BCRYPT_COST"))

	passwordEncrypted, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcryptCost)

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Password: string(passwordEncrypted),
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

	user, err := controller.userModel.Login(userRequest.Email, userRequest.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Login Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Login Succes", "token: "+user.Token))
}

// GetUserProfileController is controller for user profile
func (controller *Controller) GetUserProfileController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := controller.userModel.GetProfile(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get User Profile", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get User Profile", user))
}

// UpdatePasswordController is controller for user edit their password
func (controller *Controller) UpdatePasswordController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	var userRequest models.User
	c.Bind(&userRequest)

	if _, err := controller.userModel.UpdatePassword(userRequest, userId); err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Change Password", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Change Password", nil))
}