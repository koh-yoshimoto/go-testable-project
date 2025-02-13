package domain

import "time"

type Task struct {
	ID          int64
	Status			string
	Title       string
	Description string
	DueDate     *time.Time
}

func (t *Task) IsOverdue() bool {
	return t.DueDate.Before(time.Now())
}
