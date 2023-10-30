package report

import "time"

func NewReportFixture(id string, userID string, content string, reportingTime time.Time) *Report {
	return &Report{
		id:            id,
		userID:        userID,
		content:       content,
		reportingTime: reportingTime,
	}
}
