package routers

import (
	"bootcamp/handlers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.POST("/user", handlers.CreateUser)
	e.GET("/user/:id", handlers.GetUser)
	e.GET("/users", handlers.GetUsers)
	e.PUT("/user/:id", handlers.UpdateUser)
	e.DELETE("/user/:id", handlers.DeleteUser)
}
