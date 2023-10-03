package application

import (
	"context"
	"fmt"
	"time"

	"github.com/citcho/go-gizlog/internal/common/auth"
	"github.com/citcho/go-gizlog/internal/common/derror"
	"github.com/citcho/go-gizlog/internal/report/domain/report"
)

type IReportUsecase interface {
	StoreReport(context.Context, StoreReportCommand) error
}

type ReportUsecase struct {
	service    report.IReportService
	repository report.IReportRepository
}

func NewReportUsecase(rs report.IReportService, rr report.IReportRepository) *ReportUsecase {
	return &ReportUsecase{
		service:    rs,
		repository: rr,
	}
}

func (ru *ReportUsecase) StoreReport(ctx context.Context, cmd StoreReportCommand) (err error) {
	defer derror.Wrap(&err, "StoreReport(%+v)", cmd)

	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("user_id not found")
	}

	r, err := report.NewReport(
		cmd.ID,
		userId,
		cmd.Content,
		cmd.ReportingTime,
	)
	if err != nil {
		return fmt.Errorf("%s: %w", err, derror.InvalidArgument)
	}

	exists, err := ru.service.Exists(ctx, r)
	if err != nil {
		return fmt.Errorf("%s: %w", err, derror.InvalidArgument)
	}
	if exists {
		return fmt.Errorf("既に%sの日報を作成しています。: %w", r.ReportingTime().Format(time.DateOnly), derror.InvalidArgument)
	}

	if err := ru.repository.Save(ctx, r); err != nil {
		return err
	}

	return nil
}
