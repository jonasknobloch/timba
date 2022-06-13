package tim

import (
	"time"
)

type Task struct {
	Records   []*Record `json:"records"`
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	UpdatedAt int64     `json:"updatedAt"`
	CreatedAt int64     `json:"createdAt"`
}

func NewTask(id string, title string, updatedAt, createdAt time.Time) *Task {
	return &Task{
		Records:   make([]*Record, 0),
		ID:        id,
		Title:     title,
		UpdatedAt: updatedAt.Unix(),
		CreatedAt: createdAt.Unix(),
	}
}

func (t *Task) AddRecord(record *Record) {
	t.Records = append(t.Records, record)
}
