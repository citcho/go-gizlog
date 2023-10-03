package http

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/citcho/go-gizlog/internal/common/derror"
	"github.com/citcho/go-gizlog/internal/report/application"
	"github.com/labstack/echo/v4"
	"github.com/oklog/ulid/v2"
)

type reportController struct {
	usecase application.IReportUsecase
}

func NewReportController(ru application.IReportUsecase) *reportController {
	return &reportController{
		usecase: ru,
	}
}

func (rc reportController) storeReport(ctx echo.Context) error {
	var cmd application.StoreReportCommand
	if err := ctx.Bind(&cmd); err != nil {
		return ctx.JSON(
			http.StatusBadRequest,
			&derror.ErrResponse{Message: err.Error()},
		)
	}

	cmd.ID = ulid.Make().String()

	if err := rc.usecase.StoreReport(ctx.Request().Context(), cmd); err != nil {
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
		fmt.Sprintf("http://localhost:8080/reports/%s", cmd.ID),
	)

	return ctx.JSON(http.StatusCreated, nil)
}
