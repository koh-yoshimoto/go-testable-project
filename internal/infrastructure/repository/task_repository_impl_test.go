package repository

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/koh-yoshimoto/go-testable-project/internal/domain"
)

func TestTaskRepositoryImpl_GetByID(t *testing.T) {

	prepareTestDB(t)

	tests := map[string]struct {
		give int64
		want *domain.Task
		wantErr error
	}{
		"success": {
			give: 1,
			want: &domain.Task{
				ID: 1,
				Title: "Task 1",
				Description: "This is the first task",
				Status: "new",
				DueDate: nil,
			},
		},
		"not_found": {
			give: 9999,
			wantErr: errors.New("sql: no rows in result set"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			taskRepo := NewTaskRepository(testDB)

			got, err := taskRepo.GetByID(tt.give)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.wantErr.Error(), err.Error()) != "" {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
