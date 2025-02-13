package usecase

import "github.com/koh-yoshimoto/go-testable-project/internal/domain"

type TaskUsecase struct {
	repo domain.TaskRepository	
}

func NewTaskUsecase(repo domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{repo: repo}
}

func (u *TaskUsecase) GetTask(id int64) (*domain.Task, error) {
	return u.repo.GetByID(id)
}

func (u *TaskUsecase) CreateTask(task *domain.Task) error {
	return u.repo.Create(task)
}
