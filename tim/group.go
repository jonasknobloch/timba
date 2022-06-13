package tim

import (
	"time"
)

type Group struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	UpdatedAt int64  `json:"updatedAt"`
	CreatedAt int64  `json:"createdAt"`
}

func NewGroup(id string, title string, updatedAt, createdAt time.Time) *Group {
	return &Group{
		ID:        id,
		Title:     title,
		UpdatedAt: updatedAt.UnixMilli(),
		CreatedAt: createdAt.UnixMilli(),
	}
}
