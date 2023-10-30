package report_test

import (
	"strings"
	"testing"
	"time"

	"github.com/citcho/go-gizlog/internal/common/clock"
	"github.com/citcho/go-gizlog/internal/report/domain/report"
	"github.com/google/go-cmp/cmp"
	"github.com/oklog/ulid/v2"
)

func TestNewReport(t *testing.T) {
	uid := ulid.Make().String()
	rid := ulid.Make().String()
	c := clock.FixedClocker{}

	type args struct {
		id            string
		userId        string
		content       string
		reportingTime time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *report.Report
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				id:            rid,
				userId:        uid,
				content:       "content",
				reportingTime: c.Now(),
			},
			want: report.NewReportFixture(
				rid,
				uid,
				"content",
				c.Now(),
			),
			wantErr: false,
		},
		{
			name: "異常系: 日報ID無しで日報を作成できない",
			args: args{
				id:            "",
				userId:        uid,
				content:       "content",
				reportingTime: c.Now(),
			},
			want:    &report.Report{},
			wantErr: true,
		},
		{
			name: "異常系: ユーザーID無しで日報を作成できない",
			args: args{
				id:            rid,
				userId:        "",
				content:       "content",
				reportingTime: c.Now(),
			},
			want:    &report.Report{},
			wantErr: true,
		},
		{
			name: "正常系: contentが1000文字を超えないと日報を作成できる",
			args: args{
				id:            rid,
				userId:        uid,
				content:       strings.Repeat("a", 999),
				reportingTime: c.Now(),
			},
			want: report.NewReportFixture(
				rid,
				uid,
				strings.Repeat("a", 999),
				c.Now(),
			),
			wantErr: false,
		},
		{
			name: "異常系: contentが1000文字を超えると日報を作成できない",
			args: args{
				id:            rid,
				userId:        uid,
				content:       strings.Repeat("a", 1000),
				reportingTime: c.Now(),
			},
			want:    &report.Report{},
			wantErr: true,
		},
		{
			name: "異常系: 内容無しで日報を作成できない",
			args: args{
				id:            rid,
				userId:        uid,
				content:       "",
				reportingTime: c.Now(),
			},
			want:    &report.Report{},
			wantErr: true,
		},
		{
			name: "異常系: 明日以降の日報を作成できない",
			args: args{
				id:            rid,
				userId:        uid,
				content:       "content",
				reportingTime: c.Now().AddDate(0, 0, 1),
			},
			want:    &report.Report{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := report.NewReport(tt.args.id, tt.args.userId, tt.args.content, tt.args.reportingTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewReport() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(report.Report{})
			if d := cmp.Diff(got, tt.want, opt); len(d) != 0 {
				t.Errorf("differs: (-got +want)\n%s", d)
			}
		})
	}
}
