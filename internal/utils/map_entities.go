package utils

import (
	"errors"
	"ozonva/ova-task-api/internal/pkg/entities/tasks"
)

var ErrDuplicateKey = errors.New("an item with the same key has already been added")
var ErrArgumentNil = errors.New("value cannot be null")

func MapTasks(source []tasks.Task) (map[uint64]tasks.Task, error) {
	if source == nil {
		return nil, ErrArgumentNil
	}
	var result = make(map[uint64]tasks.Task, len(source))
	for _, value := range source {
		if _, contains := result[value.TaskId]; contains {
			return nil, ErrDuplicateKey
		}
		result[value.TaskId] = value
	}
	return result, nil
}
