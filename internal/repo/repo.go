package repo

import (
	"ozonva/ova-task-api/internal/pkg/entities/tasks"
)

type Repo interface {
	AddTasks(entities []tasks.Task) error
	ListTasks(limit, offset uint64) ([]tasks.Task, error)
	DescribeTasks(entityId uint64) (*tasks.Task, error)
}
