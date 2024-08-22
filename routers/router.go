package routers

import (
	"bootcamp/handlers"
	"bootcamp/middleware"
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo, db *sql.DB, rdb *redis.Client) {
	e.POST("/user", handlers.CreateUser(db), middleware.BasicAuthMiddleware())
	e.GET("/user/:id", handlers.GetUser(db, rdb), middleware.BasicAuthMiddleware())
	e.PUT("/user/:id", handlers.UpdateUser(db, rdb), middleware.BasicAuthMiddleware())
	e.DELETE("/user/:id", handlers.DeleteUser(db, rdb), middleware.BasicAuthMiddleware())
}
