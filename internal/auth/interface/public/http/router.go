package http

import (
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo, ac authController) {
	e.POST("/login", func(c echo.Context) error { return ac.Login(c) })
}
