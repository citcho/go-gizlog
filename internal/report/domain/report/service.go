package report

import (
	"context"
)

type IReportRepository interface {
	Exists(context.Context, *Report) (bool, error)
}

type ReportService struct {
	repository IReportRepository
}

func NewReportService(rr IReportRepository) *ReportService {
	return &ReportService{
		repository: rr,
	}
}

func (rs ReportService) Exists(ctx context.Context, r *Report) (bool, error) {
	exists, err := rs.repository.Exists(ctx, r)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
