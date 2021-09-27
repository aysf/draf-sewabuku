package main

import (
	"sewabuku/database"
	"sewabuku/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.Connect()

	e := echo.New()
	routes.Setup(e)

	e.Logger.Fatal(e.Start(":8080"))
}
