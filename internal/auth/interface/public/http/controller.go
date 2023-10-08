package http

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/citcho/go-gizlog/internal/auth/application"
	"github.com/citcho/go-gizlog/internal/common/derror"
	"github.com/labstack/echo/v4"
)

type authController struct {
	usecase application.IAuthUsecase
}

func NewAuthController(au application.IAuthUsecase) *authController {
	return &authController{
		usecase: au,
	}
}

func (ac authController) Login(ctx echo.Context) error {
	var cmd application.LoginCommand
	if err := ctx.Bind(&cmd); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			&derror.ErrResponse{Message: err.Error()},
		)
	}

	jwt, err := ac.usecase.Login(ctx.Request().Context(), cmd)
	if err != nil {
		switch {
		case errors.Is(err, derror.InvalidArgument):
			return ctx.JSON(
				http.StatusBadRequest,
				&derror.ErrResponse{Message: err.Error()},
			)
		default:
			return ctx.JSON(
				http.StatusInternalServerError,
				&derror.ErrResponse{Message: err.Error()},
			)
		}
	}

	cookie := http.Cookie{
		Name:    "token",
		Value:   jwt,
		Expires: time.Now().Add(24 * time.Hour),
		Path:    "/",
		Domain:  os.Getenv("API_DOMAIN"),
		// Secure:   true,
		// HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	}
	ctx.SetCookie(&cookie)

	return ctx.JSON(http.StatusOK, nil)
}
