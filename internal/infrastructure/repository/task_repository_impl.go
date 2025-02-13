package repository

import (
	"database/sql"
	"github.com/koh-yoshimoto/go-testable-project/internal/domain"
)

type TaskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) domain.TaskRepository {
	return &TaskRepositoryImpl{db: db}
}

func (r *TaskRepositoryImpl) GetByID(id int64) (*domain.Task, error) {
	task := &domain.Task{}
	err := r.db.QueryRow("SELECT id, title, status, due_date, description FROM tasks WHERE id = ?", id).
		Scan(&task.ID, &task.Title, &task.Status, &task.DueDate, &task.Description)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepositoryImpl) Create(task *domain.Task) error {
	_, err := r.db.Exec("INSERT INTO tasks (id, title, status, due_date, description) VALUES (?, ?, ?, ?)",
		task.Title, task.Description, task.DueDate, task.Status)
	return err
}

func (r *TaskRepositoryImpl) Update(task *domain.Task) error {
	_, err := r.db.Exec("UPDATE tasks SET title = ?, description = ?, due_date = ?, completed = ? WHERE id = ?",
		task.Title, task.Description, task.DueDate, task.Status, task.ID)
	return err
}

func (r *TaskRepositoryImpl) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

