package auth

import (
	"net/http"

	"github.com/citcho/go-gizlog/internal/common/derror"
	"github.com/labstack/echo/v4"
)

func (j *JWTer) JwtHttpMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req, err := j.FillContext(c.Request())
		if err != nil {
			return c.JSON(
				http.StatusUnauthorized,
				derror.ErrResponse{
					Message: err.Error(),
				},
			)
		}

		c.SetRequest(req)

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
