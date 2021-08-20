package repo

import (
	"ozonva/ova-task-api/internal/pkg/entities/tasks"
)

type Repo interface {
	AddEntities(entities []tasks.Task) error
	ListEntities(limit, offset uint64) ([]tasks.Task, error)
	DescribeEntity(entityId uint64) (*tasks.Task, error)
}
