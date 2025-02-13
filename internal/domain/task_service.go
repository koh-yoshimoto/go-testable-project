package domain

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) MarkComplete(id int64) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	task.Status = "completed"
	return s.repo.Update(task)
}
