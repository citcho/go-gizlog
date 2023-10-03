package report

import (
	"errors"
	"time"
	"unicode/utf8"
)

type Report struct {
	id            string
	userID        string
	content       string
	reportingTime time.Time
}

func NewReport(
	id string,
	userId string,
	content string,
	reportingTime time.Time,
) (*Report, error) {
	if id == "" {
		return &Report{}, errors.New("idがありません。")
	}
	if userId == "" {
		return &Report{}, errors.New("user_idがありません。")
	}
	if utf8.RuneCountInString(content) > 1000 {
		return &Report{}, errors.New("投稿できる日報内容の文字数を超過しています。")
	}
	if content == "" {
		return &Report{}, errors.New("日報内容がありません。")
	}
	if reportingTime.After(time.Now().AddDate(0, 0, 1)) {
		return &Report{}, errors.New("明日以降の日報は作成できません。")
	}

	r := &Report{
		id:            id,
		userID:        userId,
		content:       content,
		reportingTime: reportingTime,
	}

	return r, nil
}

func (r Report) ID() string {
	return r.id
}

func (r Report) UserID() string {
	return r.userID
}

func (r Report) Content() string {
	return r.content
}

func (r Report) ReportingTime() time.Time {
	return r.reportingTime
}