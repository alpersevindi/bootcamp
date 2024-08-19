package main

import (
	"bootcamp/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routers.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
