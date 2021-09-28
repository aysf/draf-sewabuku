package user

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"sewabuku/models"
	"sewabuku/util"
	"strconv"

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
		return c.JSON(http.StatusInternalServerError, util.ResponseFail("Register Fail", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Register Success", nil))
}

func (controller *Controller) LoginUserController(c echo.Context) error {
	var userRequest models.User
	c.Bind(&userRequest)

	user, err := controller.userModel.Login(userRequest.Email, userRequest.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Login Failed", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Login Succes", "token: "+user.Token))
}

func (controller *Controller) GetUserProfileController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	user, err := controller.userModel.GetProfile(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Get User Profile", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Get User Profile", user))
}

func (controller *Controller) UpdatePasswordController(c echo.Context) error {
	userId := middlewares.ExtractTokenUserId(c)

	var userRequest models.User
	c.Bind(&userRequest)

	if _, err := controller.userModel.UpdatePassword(userRequest, userId); err != nil {
		return c.JSON(http.StatusBadRequest, util.ResponseFail("Fail to Change Password", nil))
	}

	return c.JSON(http.StatusOK, util.ResponseSuccess("Success Change Password", nil))
}