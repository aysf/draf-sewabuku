package controllers

import (
	"net/http"
	"os"
	"sewabuku/database"
	"sewabuku/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = os.Getenv("SECRET_KEY")

func Register(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)

	user := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: string(password),
	}

	database.DB.Create(&user)

	return c.JSON(http.StatusAccepted, user)
}

func Login(c echo.Context) error {
	userInput := new(models.User)

	if err := c.Bind(userInput); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", userInput.Email).First(&user)

	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "email not exist",
		})

	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "could not login",
		})
	}

	cookie := http.Cookie{
		Name:    "jwt",
		Value:   token,
		Expires: time.Now().Add(time.Hour * 24),
	}

	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}

func User(c echo.Context) error {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		return err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorized",
		})

	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

func Logout(c echo.Context) error {
	cookie := http.Cookie{
		Name:    "jwt",
		Value:   "",
		Expires: time.Now().Add(-time.Hour),
	}

	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "success",
	})
}
