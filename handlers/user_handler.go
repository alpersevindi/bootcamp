package handlers

import (
	"bootcamp/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var users = map[uuid.UUID]models.User{}

func CreateUser(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.UUID = uuid.New()
	users[user.UUID] = *user
	return c.JSON(http.StatusCreated, user)
}

func GetUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}

	user, exists := users[id]
	if !exists {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

func GetUsers(c echo.Context) error {
	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, "There is no user")
	}
	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}
	if _, exists := users[id]; !exists {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	user.UUID = id
	users[id] = *user
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid UUID")
	}
	if _, exists := users[id]; !exists {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}
