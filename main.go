package main

import (
	"bootcamp/database"
	"bootcamp/routers"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.InitMySQL()
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	routers.InitRoutes(e, db)
	e.Logger.Fatal(e.Start(":8080"))
}
