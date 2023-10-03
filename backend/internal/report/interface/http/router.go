package http

import (
	"github.com/citcho/go-gizlog/internal/common/auth"
	"github.com/labstack/echo/v4"
)

func AddRoutes(e *echo.Echo, j *auth.JWTer, rc reportController) {
	g := e.Group("/reports")
	g.Use(j.JwtHttpMiddleware)
	g.POST("", func(c echo.Context) error { return rc.storeReport(c) })
}
