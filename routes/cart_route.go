package routes

import (
	"os"
	"sewabuku/controllers/cart"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CartPath(e *echo.Echo, cartController *cart.Controller) {
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))

	// rent book
	jwtAuth.POST("/cart/rent", cartController.RentBook)

	// return book
	jwtAuth.PUT("/cart/return", cartController.ReturnBook)

	// list book
	jwtAuth.GET("/cart", cartController.ListBook)

}
