package tasks

import (
	"fmt"
	"ozonva/ova-task-api/pkg/entities/tasks/task_states"
	"time"
)

type Task struct {
	UserId      uint64
	TaskId      uint64
	Description string
	DateCreated time.Time
	State       task_states.TaskState
}

func New(userId uint64, taskId uint64, description string, dateCreated time.Time) *Task {
	return &Task{
		UserId:      userId,
		TaskId:      taskId,
		Description: description,
		DateCreated: dateCreated,
	}
}

func (task *Task) String() string {
	return fmt.Sprintf("[User:%v, task:%v] %v - %v", task.UserId, task.TaskId, task.DateCreated.Format("2006.01.02 15:04:05"), task.Description)
}
