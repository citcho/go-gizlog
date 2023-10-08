package application_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/citcho/go-gizlog/internal/common/config"
	"github.com/citcho/go-gizlog/internal/common/database"
	"github.com/citcho/go-gizlog/internal/report/application"
	"github.com/citcho/go-gizlog/internal/report/infrastructure/report"
	"github.com/golang/mock/gomock"
)

func Test_ReportUsecase_StoreReport(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.NewDBConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	db := database.NewDB(*cfg)

	repository := report.NewMySQLRepository(db)
	sut := application.NewReportUsecase(repository)

	type args struct {
		ctx context.Context
		cmd application.StoreReportCommand
	}

	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "今日の日付で日報を作成できること",
			args: args{
				ctx: context.Background(),
				cmd: application.StoreReportCommand{
					ID:            "dummy-uuid-string",
					Content:       "content",
					ReportingTime: time.Now(),
				},
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := sut.StoreReport(tt.args.ctx, tt.args.cmd)
			if err != nil {
				t.Errorf("ReportUsecase.StoreReport() error = %v", err)
			}
		})
	}
}
