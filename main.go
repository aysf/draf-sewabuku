package main

import (
	"sewabuku/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	routes.Setup(e)

	e.Logger.Fatal(e.Start(":8080"))
}
