//go:build integration

package application_test

import (
	"context"
	"testing"
	"time"

	"github.com/citcho/go-gizlog/internal/common/auth"
	"github.com/citcho/go-gizlog/internal/common/clock"
	"github.com/citcho/go-gizlog/internal/common/config"
	"github.com/citcho/go-gizlog/internal/common/database"
	"github.com/citcho/go-gizlog/internal/common/testutil"
	"github.com/citcho/go-gizlog/internal/report/application"
	"github.com/citcho/go-gizlog/internal/report/domain/report"
	"github.com/citcho/go-gizlog/internal/report/infrastructure/dao"
	report_dao "github.com/citcho/go-gizlog/internal/report/infrastructure/dao"
	report_infra "github.com/citcho/go-gizlog/internal/report/infrastructure/report"
	"github.com/google/go-cmp/cmp"
	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"
)

func TestReportUsecase_StoreReport(t *testing.T) {
	t.Helper()

	// Arrange
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTx(ctx, nil)
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := config.NewDBConfig()
	if err != nil {
		t.Fatal(err)
	}
	db := database.NewDB(cfg)
	repo := report_infra.NewMySQLRepository(db)
	svc := report.NewReportService(repo)

	uid := ulid.Make().String()
	rid := ulid.Make().String()
	c := clock.FixedClocker{}

	cmd := application.StoreReportCommand{
		ID:            rid,
		Content:       "test content",
		ReportingTime: c.Now(),
	}

	want, err := report.NewReport(
		rid,
		uid,
		"test content",
		c.Now(),
	)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

	ctx = auth.SetUserID(ctx, uid)

	// Act
	sut := application.NewReportUsecase(svc, repo)
	if err = sut.StoreReport(ctx, cmd); err != nil {
		t.Fatalf("%s", err.Error())
	}

	// Assert
	got, err := findReportById(ctx, db, cmd.ID)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	opt := cmp.AllowUnexported(report.Report{})
	if d := cmp.Diff(got, want, opt); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}

func findReportById(ctx context.Context, db *bun.DB, reportId string) (*report.Report, error) {
	var r report_dao.Report
	err := db.NewSelect().
		Model(&r).
		Where("id = ?", reportId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	report := report.ReConstructFromRepository(
		r.ID,
		r.UserID,
		r.Content,
		r.ReportingTime,
	)

	return report, nil
}

func prepareReport(rid string, uid string, content string, reportingTime time.Time) error {
	cfg, err := config.NewDBConfig()
	if err != nil {
		return err
	}
	db := database.NewDB(cfg)
	dao := dao.Report{
		ID:            rid,
		UserID:        uid,
		Content:       content,
		ReportingTime: reportingTime,
	}

	db.NewInsert().
		Model(&dao).
		Exec(context.Background())

	return nil
}
