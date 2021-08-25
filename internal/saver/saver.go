package saver

import (
	"fmt"
	"ozonva/ova-task-api/internal/flusher"
	"ozonva/ova-task-api/internal/pkg/entities/tasks"
	"sync"
	"time"
)

type Saver interface {
	Save(entity tasks.Task)
	Close()
}

type saver struct {
	buffer chan tasks.Task
	wait   *sync.WaitGroup
}

func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	savePeriod time.Duration,
) Saver {
	if savePeriod <= 0 {
		panic("savePeriod must be greater than 0")
	}
	if flusher == nil {
		panic("flusher is nil")
	}
	var buffer = make(chan tasks.Task, capacity)
	saver := &saver{}
	saver.init(buffer, flusher, savePeriod)
	return saver
}

func (s *saver) Save(task tasks.Task) {
	s.buffer <- task
}

func (s *saver) Close() {
	close(s.buffer)
	s.wait.Wait()
}

func (s *saver) init(buffer chan tasks.Task, flusher flusher.Flusher, savePeriod time.Duration) {
	var flush = func(toFlush []tasks.Task) []tasks.Task {
		if len(toFlush) > 0 {
			notFlushed := flusher.Flush(toFlush)
			return notFlushed
		}
		return toFlush
	}

	s.buffer = buffer
	s.wait = &sync.WaitGroup{}
	go func() {
		s.wait.Add(1)
		defer s.wait.Done()

		toFlush := make([]tasks.Task, 0, 1)
		for {
			select {
			case entity, ok := <-buffer:
				if !ok {
					toFlush = flush(toFlush)
					if len(toFlush) > 0 {
						fmt.Printf("%v tasks were not saved\n", len(toFlush))
					}
					return
				}
				toFlush = append(toFlush, entity)
			case <-time.After(savePeriod):
				toFlush = flush(toFlush)
				break
			}
		}
	}()
}
