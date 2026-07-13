package models

import "time"


type Task struct {
	Title string
	Description string
	IsComplete bool
	CreatedAt time.Time
	CompletedAt *time.Time
}