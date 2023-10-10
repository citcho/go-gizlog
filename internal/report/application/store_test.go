//go:build integration

package application_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/citcho/go-gizlog/internal/common/auth"
	"github.com/citcho/go-gizlog/internal/common/clock"
	"github.com/citcho/go-gizlog/internal/common/config"
	"github.com/citcho/go-gizlog/internal/common/database"
	"github.com/citcho/go-gizlog/internal/report/application"
	"github.com/citcho/go-gizlog/internal/report/domain/report"
	report_dao "github.com/citcho/go-gizlog/internal/report/infrastructure/dao"
	report_infra "github.com/citcho/go-gizlog/internal/report/infrastructure/report"
	user_dao "github.com/citcho/go-gizlog/internal/user/infrastructure/dao"
	"github.com/google/go-cmp/cmp"
	"github.com/oklog/ulid/v2"
	"github.com/uptrace/bun"
)

func TestReportUsecase_StoreReport(t *testing.T) {
	t.Helper()

	// Arrange
	cfg, err := config.NewDBConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	db := database.NewDB(cfg)
	repo := report_infra.NewMySQLRepository(db)
	svc := report.NewReportService(repo)

	uid := ulid.Make().String()
	ctx := auth.SetUserID(context.Background(), uid)
	prepareUser(ctx, db, "testuser", "test@example.com", "P@ssw0rd")

	rid := ulid.Make().String()
	c := clock.FixedClocker{}
	cmd := application.StoreReportCommand{
		ID:            rid,
		Content:       "test-content",
		ReportingTime: c.Now(),
	}
	want, err := report.NewReport(
		rid,
		uid,
		"test-content",
		time.Date(2022, 5, 10, 0, 0, 0, 0, time.UTC),
	)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}

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

func prepareUser(ctx context.Context, db *bun.DB, name string, email string, password string) error {
	userId, ok := auth.GetUserID(ctx)
	if !ok {
		return fmt.Errorf("user_id not found")
	}

	_, err := db.NewInsert().
		Model(&user_dao.User{
			ID:       userId,
			Name:     name,
			Email:    email,
			Password: password,
		}).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
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
