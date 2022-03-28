package contollers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetUserById(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func UpdateUserById(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func CreateUser(c echo.Context) error {
	return c.JSON(http.StatusCreated, "")
}
