package routes

import (
	"github.com/labstack/echo/v4"
	"shorty/contollers"
)

func DefineApiEndpoints(s *echo.Echo) {
	g := s.Group("/api")

	// User
	g.GET("/user/:id", contollers.GetUserById)
	g.PUT("/user/:id", contollers.UpdateUserById)
	g.POST("/user", contollers.CreateUser)

	// Links
	g.GET("/links/:id", contollers.GetLinkById)
	g.PUT("/links/:id", contollers.UpdateLinkById)
	g.POST("/links", contollers.CreateLink)
}
