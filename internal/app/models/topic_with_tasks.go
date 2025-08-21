package models

import "time"

type TopicWithTasks struct {
	ID    int64
	Name  string
	Tasks []Task

	CreatedAt time.Time
	UpdatedAt time.Time
}
