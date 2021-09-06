package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"ozonva/ova-task-api/internal/metrics"
	repopkg "ozonva/ova-task-api/internal/repo"
	"ozonva/ova-task-api/internal/utils"
	desc "ozonva/ova-task-api/pkg/api/ova-task-api"
	. "ozonva/ova-task-api/pkg/entities/tasks"
	"time"
)

type OvaTaskAPI struct {
	desc.UnimplementedOvaTaskApiServer
	repo      repopkg.Repo
	batchSize int
}

func NewOvaTaskApi(repo repopkg.Repo) desc.OvaTaskApiServer {
	return &OvaTaskAPI{
		repo: repo,
		// todo pass as argument
		batchSize: 2,
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

	task := New(in.GetTaskTemplate().GetUserId(), 0, in.GetTaskTemplate().GetDescription(), time.Now())
	err := o.repo.AddTasks(ctx, []Task{*task})
	if err != nil {
		return nil, err
	}

	metrics.Inc(metrics.CreateTaskCounter)
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

	metrics.Inc(metrics.RemoveTaskCounter)
	return &emptypb.Empty{}, nil
}

func (o OvaTaskAPI) MultiCreateTaskV1(ctx context.Context, in *desc.MultiCreateTaskV1Request) (*desc.MultiCreateTaskV1Response, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "MultiCreateTaskV1")
	span.SetTag("tasks", len(in.TaskTemplate))
	defer span.Finish()

	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("CreateTaskV1")

	taskTemplates := in.GetTaskTemplate()
	tasksToAdd := make([]Task, 0, len(taskTemplates))
	for _, inTask := range taskTemplates {
		task := New(inTask.GetUserId(), 0, inTask.GetDescription(), time.Now())
		tasksToAdd = append(tasksToAdd, *task)
	}

	for _, tasksSlice := range utils.SplitTasksSlice(tasksToAdd, o.batchSize) {
		childSpan, ctx := opentracing.StartSpanFromContext(ctx, "add chunk of tasks")
		childSpan.SetTag("tasks_chunk", len(tasksSlice))
		err := o.repo.AddTasks(ctx, tasksSlice)
		if err != nil {
			log.Error().Err(err).Send()
			return nil, err
		}
		childSpan.Finish()
	}
	return &desc.MultiCreateTaskV1Response{}, nil
}

func (o OvaTaskAPI) UpdateTaskV1(ctx context.Context, in *desc.UpdateTaskV1Request) (*desc.UpdateTaskV1Response, error) {
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("CreateTaskV1")

	task := New(in.GetUserId(), in.GetTaskId(), in.GetDescription(), time.Time{})
	err := o.repo.UpdateTask(ctx, *task)
	if err != nil {
		return nil, err
	}

	metrics.Inc(metrics.UpdateTaskCounter)
	return &desc.UpdateTaskV1Response{}, nil
}
