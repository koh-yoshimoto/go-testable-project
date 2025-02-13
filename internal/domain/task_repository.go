package domain

//go:generate mockgen -source=task_repository.go -destination=../mocks/mock_task_repository.go -package=mocks
type TaskRepository interface {
	GetByID(id int64) (*Task, error)
	Create(task *Task) error
	Update(task *Task) error
	Delete(id int64) error
}
