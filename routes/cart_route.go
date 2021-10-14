package routes

import (
	"os"
	"sewabuku/controllers/cart"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CartPath(e *echo.Echo, cartController *cart.Controller) {
	jwtAuth := e.Group("/carts")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("SECRET_KEY"))))

	// rent book
	jwtAuth.POST("/rent", cartController.RentBook)

	// return book
	jwtAuth.PUT("/return", cartController.ReturnBook)

	// extend book
	jwtAuth.PUT("/extend", cartController.ExtendDateDue)

	// list book
	jwtAuth.GET("/", cartController.ListBook)

}
