package contollers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetLinkById(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func UpdateLinkById(c echo.Context) error {
	return c.JSON(http.StatusOK, "")
}

func CreateLink(c echo.Context) error {
	return c.JSON(http.StatusCreated, "")
}

func LinkRedirect(c echo.Context) error {
	return c.String(http.StatusMovedPermanently, c.Param("hash"))
}
