package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/koh-yoshimoto/go-testable-project/internal/domain"
	"github.com/koh-yoshimoto/go-testable-project/internal/mocks"
)

func TestTaskUsecase_GetTask(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		give  int
		want  *domain.Task
		setup func(ctrl *gomock.Controller) *TaskUsecase
	}{
		"success": {
			give: 1,
			want: &domain.Task{ID: 1, Title: "test", Status: "active"},
			setup: func(ctrl *gomock.Controller) *TaskUsecase {
				taskRepo := mocks.NewMockTaskRepository(ctrl)
				task := &domain.Task{ID: 1, Title: "test", Status: "active"}
				taskRepo.EXPECT().GetByID(gomock.Any()).Return(task, nil)
				return NewTaskUsecase(taskRepo)
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			taskUsecase := tt.setup(ctrl)
			got, _ := taskUsecase.GetTask(int64(tt.give))
			if cmp.Diff(got, tt.want) != "" {
				t.Errorf("want: %v, got: %v", tt.want, got)
			}
		})
	}
}
