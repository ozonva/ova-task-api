package utils

import (
	"math"
	"ozonva/ova-task-api/pkg/entities/tasks"
)

func SplitSlice(source []string, batch int) [][]string {
	if batch <= 0 {
		panic("Batch must be greater than 0.")
	}
	chunksCount := int(math.Ceil(float64(len(source)) / float64(batch)))
	if chunksCount == 0 {
		return [][]string{}
	}

	chunks := make([][]string, 0, chunksCount)
	for i := 0; i < len(source); i += batch {
		from := i
		to := from + batch
		if to > len(source) {
			to = len(source)
		}
		chunks = append(chunks, source[from:to])
	}
	return chunks
}

func SplitTasksSlice(entities []tasks.Task, batchSize int) [][]tasks.Task {
	if batchSize <= 0 {
		panic("Batch must be greater than 0.")
	}
	chunksCount := int(math.Ceil(float64(len(entities)) / float64(batchSize)))
	if chunksCount == 0 {
		return [][]tasks.Task{}
	}

	chunks := make([][]tasks.Task, 0, chunksCount)
	for i := 0; i < len(entities); i += batchSize {
		from := i
		to := from + batchSize
		if to > len(entities) {
			to = len(entities)
		}
		chunks = append(chunks, entities[from:to])
	}
	return chunks
}
