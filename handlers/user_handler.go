package handlers

import (
	"bootcamp/models"
	"database/sql"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var users = map[uuid.UUID]models.User{}

func CreateUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user.UUID = uuid.New()

		_, err := db.Exec("INSERT INTO users (uuid, name, surname) VALUES (?, ?, ?)", user.UUID, user.Name, user.Surname)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, user)
	}
}

func GetUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid UUID")
		}

		var user models.User
		err = db.QueryRow("SELECT uuid, name, surname FROM users WHERE uuid = ?", id).
			Scan(&user.UUID, &user.Name, &user.Surname)

		if err != nil {
			if err == sql.ErrNoRows {
				return c.JSON(http.StatusNotFound, "User not found")
			}
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	}
}

func GetUsers(c echo.Context) error {
	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, "There is no user")
	}
	return c.JSON(http.StatusOK, users)
}

func UpdateUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid UUID")
		}

		user := new(models.User)
		if err := c.Bind(user); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		user.UUID = id

		_, err = db.Exec("UPDATE users SET name = ?, surname = ? WHERE uuid = ?", user.Name, user.Surname, user.UUID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid UUID")
		}

		_, err = db.Exec("DELETE FROM users WHERE uuid = ?", id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
