package report

import (
	"context"
	"time"

	"github.com/citcho/go-gizlog/internal/report/domain/report"
	"github.com/citcho/go-gizlog/internal/report/infrastructure/dao"
	"github.com/uptrace/bun"
)

type MySQLRepository struct {
	db *bun.DB
}

func NewMySQLRepository(db *bun.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (mr *MySQLRepository) Save(ctx context.Context, report *report.Report) error {
	dao := dao.Report{
		ID:            report.ID(),
		UserID:        report.UserID(),
		Content:       report.Content(),
		ReportingTime: report.ReportingTime(),
	}

	_, err := mr.db.NewInsert().
		Model(&dao).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (mr *MySQLRepository) Exists(ctx context.Context, report *report.Report) (bool, error) {
	exists, err := mr.db.NewSelect().
		Model((*dao.Report)(nil)).
		Where("reporting_time = ?", report.ReportingTime().Format(time.DateOnly)).
		Exists(ctx)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
