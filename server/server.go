package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"shorty/routes"
)

func New() *echo.Echo {

	// Server instance
	server := echo.New()

	// Define routes
	routes.DefineApiEndpoints(server)
	routes.DefineWebEndpoints(server)

	// Middlewares
	server.Use(middleware.Logger())

	return server
}
