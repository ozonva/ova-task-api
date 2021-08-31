package repo

import (
	"ozonva/ova-task-api/pkg/entities/tasks"
)

type Repo interface {
	AddTasks(tasks []tasks.Task) error
	ListTasks(limit, offset uint64) ([]tasks.Task, error)
	DescribeTasks(taskId uint64) (*tasks.Task, error)
}

type repo struct{}

func NewRepo() Repo {
	return &repo{}
}

func (*repo) AddTasks(tasks []tasks.Task) error {
	panic("implement me")
}

func (*repo) ListTasks(limit, offset uint64) ([]tasks.Task, error) {
	panic("implement me")
}

func (*repo) DescribeTasks(entityId uint64) (*tasks.Task, error) {
	panic("implement me")
}
