package application

import (
	"time"
)

type StoreReportCommand struct {
	ID            string
	Content       string    `json:"content"`
	ReportingTime time.Time `json:"reporting_time"`
}
