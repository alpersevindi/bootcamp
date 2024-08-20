package routers

import (
	"bootcamp/handlers"
	"database/sql"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *sql.DB) {
	e.POST("/user", handlers.CreateUser(db))
	e.GET("/user/:id", handlers.GetUser(db))
	e.GET("/users", handlers.GetUsers)
	e.PUT("/user/:id", handlers.UpdateUser(db))
	e.DELETE("/user/:id", handlers.DeleteUser(db))
}
