package routes

import (
	"github.com/labstack/echo/v4"
	"shorty/contollers"
)

func DefineWebEndpoints(s *echo.Echo) {
	// Redirect
	s.GET("/:hash", contollers.LinkRedirect)
}
