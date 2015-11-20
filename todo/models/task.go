package models

import (
	"time"
)

type Task struct {
	Id           string    `json:"id"`
	IsFinished   bool      `json:"is_finished"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	CreateTime   time.Time `json:"create_time"`
	UpdateTime   time.Time `json:"update_time"`
	FinishedTime time.Time `json:"finished_time"`
}
