package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	repopkg "ozonva/ova-task-api/internal/repo"
	desc "ozonva/ova-task-api/pkg/api/ova-task-api"
	. "ozonva/ova-task-api/pkg/entities/tasks"
	"time"
)

type OvaTaskAPI struct {
	desc.UnimplementedOvaTaskApiServer
	repo repopkg.Repo
}

func NewOvaTaskApi(repo repopkg.Repo) desc.OvaTaskApiServer {
	return &OvaTaskAPI{
		repo: repo,
	}
}

func mapTasksToModels(tasks []Task) []*desc.TaskV1 {
	models := make([]*desc.TaskV1, 0, len(tasks))
	for _, task := range tasks {
		models = append(models, mapTaskToModel(&task))
	}
	return models
}

func mapTaskToModel(task *Task) *desc.TaskV1 {
	return &desc.TaskV1{
		UserId:      task.UserId,
		TaskId:      task.TaskId,
		Description: task.Description,
		DateCreated: timestamppb.New(task.DateCreated),
	}
}

func (o OvaTaskAPI) CreateTaskV1(ctx context.Context, in *desc.CreateTaskV1Request) (*desc.CreateTaskV1Response, error) {
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("CreateTaskV1")

	task := New(in.GetUserId(), 0, in.Description, time.Now())
	err := o.repo.AddTasks(ctx, []Task{*task})
	if err != nil {
		return nil, err
	}
	return &desc.CreateTaskV1Response{}, nil
}

func (o OvaTaskAPI) DescribeTaskV1(ctx context.Context, in *desc.DescribeTaskV1Request) (*desc.DescribeTaskV1Response, error) {
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("DescribeTaskV1")

	task, err := o.repo.DescribeTask(ctx, in.GetTaskId())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Error(codes.NotFound, "task not found")
		}
		return nil, err
	}
	response := &desc.DescribeTaskV1Response{
		Task: mapTaskToModel(task),
	}
	return response, nil
}

func (o OvaTaskAPI) ListTasksV1(ctx context.Context, in *desc.ListTasksV1Request) (*desc.ListTasksV1Response, error) {
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("ListTasksV1")

	tasks, err := o.repo.ListTasks(ctx, in.GetLimit(), in.GetOffset())
	if err != nil {
		return nil, err
	}

	response := &desc.ListTasksV1Response{
		Tasks: mapTasksToModels(tasks),
	}
	return response, nil

}

func (o OvaTaskAPI) RemoveTasksV1(ctx context.Context, in *desc.RemoveTaskV1Request) (*emptypb.Empty, error) {
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("RemoveTasksV1")
	err := o.repo.RemoveTask(ctx, in.TaskId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
