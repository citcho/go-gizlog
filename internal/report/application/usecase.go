package application

import (
	"context"

	"github.com/citcho/go-gizlog/internal/report/domain/report"
)

type IReportRepository interface {
	Save(context.Context, *report.Report) error
}

type IReportService interface {
	Exists(context.Context, *report.Report) (bool, error)
}

type ReportUsecase struct {
	service    IReportService
	repository IReportRepository
}

func NewReportUsecase(rs IReportService, rr IReportRepository) *ReportUsecase {
	return &ReportUsecase{
		service:    rs,
		repository: rr,
	}
}
