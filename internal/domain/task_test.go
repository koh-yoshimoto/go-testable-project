package domain

import (
	"testing"
	"time"

	

)

func TestTask_IsOverdue(t *testing.T) {
	t.Parallel()

	toPointerTime := func(t time.Time) *time.Time {
		return &t
	}

	tests := []struct {
		name string
		task *Task
		want bool
	}{
		{
			name: "Task is overdue",
			task: &Task{DueDate: toPointerTime(time.Now().Add(-1 * time.Hour))},
			want: true,
		},
		{
			name: "Task is not overdue",
			task: &Task{DueDate: toPointerTime(time.Now().Add(1 * time.Hour))},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.task.IsOverdue(); got != tt.want {
				t.Errorf("Task.IsOverdue() = %v, want %v", got, tt.want)
			}
		})
	}
}
