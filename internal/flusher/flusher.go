package flusher

import (
	"fmt"
	taskspkg "ozonva/ova-task-api/internal/pkg/entities/tasks"
	"ozonva/ova-task-api/internal/repo"
	"ozonva/ova-task-api/internal/utils"
)

type Flusher interface {
	Flush(entities []taskspkg.Task) []taskspkg.Task
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

func (flusher *flusher) Flush(tasks []taskspkg.Task) (notFlushed []taskspkg.Task) {
	notFlushed = make([]taskspkg.Task, 0)
	for _, chunk := range utils.SplitTasksSlice(tasks, flusher.chunkSize) {
		err := flusher.entityRepo.AddTasks(chunk)
		if err != nil {
			fmt.Println("chunk not flushed", err)
			notFlushed = append(notFlushed, chunk...)
		}
	}
	return notFlushed
}
