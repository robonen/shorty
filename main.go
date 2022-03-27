package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	server := echo.New()

	// Middleware
	server.Use(middleware.Logger())

	// Routing
	server.GET("/links", addNewLink)

	server.Logger.Fatal(server.Start(":80"))
}

func addNewLink(context echo.Context) error {
	return context.Redirect(http.StatusFound, "https://google.com")
}
