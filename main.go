package main

import (
	"bootcamp/database"
	"bootcamp/models"
	"bootcamp/routers"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := database.InitMySQL()
	if err != nil {
		log.Fatal(err)
	}

	rdb := database.InitRedis()

	e := echo.New()
	e.Validator = &models.CustomValidator{Validator: validator.New()}
	routers.InitRoutes(e, db, rdb)
	e.Logger.Fatal(e.Start(":8080"))
}
