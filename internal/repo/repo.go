package repo

import (
	. "ozonva/ova-task-api/pkg/entities/tasks"
)

type Repo interface {
	AddTasks(tasks []Task) error
	ListTasks(limit, offset uint64) ([]Task, error)
	DescribeTasks(taskId uint64) (*Task, error)
	RemoveTask(taskId uint64) error
}

type repo struct{}

func NewRepo() Repo {
	return &repo{}
}

func (*repo) RemoveTask(taskId uint64) error {
	panic("implement me")
}

func (*repo) AddTasks(tasks []Task) error {
	panic("implement me")
}

func (*repo) ListTasks(limit, offset uint64) ([]Task, error) {
	panic("implement me")
}

func (*repo) DescribeTasks(entityId uint64) (*Task, error) {
	panic("implement me")
}
