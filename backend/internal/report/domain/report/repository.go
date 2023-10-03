package report

import (
	"context"
	"errors"
)

var ErrNotFound = errors.New("report not found")

type IReportRepository interface {
	Save(context.Context, *Report) error
	Exists(context.Context, *Report) (bool, error)
}
