package http

import (
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo, rc userController) {
	e.POST("/users", func(c echo.Context) error { return rc.StoreUser(c) })
}
