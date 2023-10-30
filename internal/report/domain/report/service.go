package report

import (
	"context"
	"fmt"
	"time"
)

//go:generate mockgen -source=./service.go -destination=./mock/service.go
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
	if exists {
		return exists, fmt.Errorf("既に%sの日報が存在します。", r.ReportingTime().Format(time.DateOnly))
	}

	return exists, nil
}
