package report_test

import (
	"context"
	"testing"

	"github.com/citcho/go-gizlog/internal/common/clock"
	"github.com/citcho/go-gizlog/internal/report/domain/report"
	mock_report "github.com/citcho/go-gizlog/internal/report/domain/report/mock"
	"github.com/golang/mock/gomock"
	"github.com/oklog/ulid/v2"
)

func TestReportService_Exists(t *testing.T) {
	t.Helper()
	c := clock.FixedClocker{}
	r, err := report.NewReport(ulid.Make().String(), ulid.Make().String(), "content", c.Now())
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		prepareMockFn func(*mock_report.MockIReportRepository)
	}
	type args struct {
		ctx context.Context
		r   *report.Report
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		{
			"日報が存在することを確認できる",
			fields{
				prepareMockFn: func(mrr *mock_report.MockIReportRepository) {
					mrr.EXPECT().
						Exists(context.Background(), r).
						Return(true, nil).
						Times(1)
				},
			},
			args{
				ctx: context.Background(),
				r:   r,
			},
			true,
			true,
		},
		{
			"日報が存在しないことを確認できる",
			fields{
				prepareMockFn: func(mrr *mock_report.MockIReportRepository) {
					mrr.EXPECT().
						Exists(context.Background(), r).
						Return(false, nil).
						Times(1)
				},
			},
			args{
				ctx: context.Background(),
				r:   r,
			},
			false,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mr := mock_report.NewMockIReportRepository(ctrl)
			tt.fields.prepareMockFn(mr)

			sut := report.NewReportService(mr)
			got, err := sut.Exists(tt.args.ctx, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReportService.Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReportService.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
