package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/citcho/go-gizlog/internal/common/derror"
	"github.com/citcho/go-gizlog/internal/user/application"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type IUserUsecase interface {
	StoreUser(context.Context, application.StoreUserCommand) error
}

type userController struct {
	usecase IUserUsecase
}

func NewUserController(uu IUserUsecase) *userController {
	return &userController{
		usecase: uu,
	}
}

func (uc userController) StoreUser(ctx echo.Context) error {
	var cmd application.StoreUserCommand
	if err := ctx.Bind(&cmd); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			&derror.ErrResponse{Message: err.Error()},
		)
	}

	cmd.ID = ulid.Make().String()

	if err := uc.usecase.StoreUser(ctx.Request().Context(), cmd); err != nil {
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

	ctx.Response().Header().Set(
		echo.HeaderLocation,
		fmt.Sprintf("http://localhost:8080/users/%s", cmd.ID),
	)

	return ctx.JSON(http.StatusCreated, nil)
}
