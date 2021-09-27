package controllers

import (
	"net/http"
	"sewabuku/database"
	"sewabuku/models"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	users := new([]models.User)

	if err := database.DB.Find(&users).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})

}
