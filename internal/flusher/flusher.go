package flusher

import (
	"fmt"
	"ozonva/ova-task-api/internal/pkg/entities/tasks"
	"ozonva/ova-task-api/internal/repo"
	"ozonva/ova-task-api/internal/utils"
)

type Flusher interface {
	Flush(entities []tasks.Task) []tasks.Task
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func New(
	chunkSize int,
	entityRepo repo.Repo,
) Flusher {
	if entityRepo == nil {
		panic("entityRepo is nil")
	}
	if chunkSize <= 0 {
		panic("chunkSize must be greater than 0")
	}
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

func (flusher *flusher) Flush(entities []tasks.Task) (notFlushed []tasks.Task) {
	notFlushed = make([]tasks.Task, 0, 0)
	for _, chunk := range utils.SplitTasksSlice(entities, flusher.chunkSize) {
		err := flusher.entityRepo.AddEntities(chunk)
		if err != nil {
			fmt.Println("chunk not flushed", err)
			notFlushed = append(notFlushed, chunk...)
		}
	}
	return notFlushed
}
