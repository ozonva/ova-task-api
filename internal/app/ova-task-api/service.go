package api

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	desc "ozonva/ova-task-api/pkg/api/ova-task-api"
)

type OvaTaskAPI struct {
	desc.UnimplementedOvaTaskApiServer
}

func NewOvaTaskApi() desc.OvaTaskApiServer {
	return &OvaTaskAPI{}
}

func (o OvaTaskAPI) CreateTaskV1(ctx context.Context, in *desc.CreateTaskV1Request) (*desc.CreateTaskV1Response, error) {
	fmt.Println("CreateTaskV1", in)
	return &desc.CreateTaskV1Response{}, nil
}

func (o OvaTaskAPI) DescribeTaskV1(ctx context.Context, in *desc.DescribeTaskV1Request) (*desc.DescribeTaskV1Response, error) {
	fmt.Println("DescribeTaskV1", in)
	return &desc.DescribeTaskV1Response{}, nil
}

func (o OvaTaskAPI) ListTasksV1(ctx context.Context, in *desc.ListTasksV1Request) (*desc.ListTasksV1Response, error) {
	fmt.Println("ListTasksV1", in)
	return &desc.ListTasksV1Response{}, nil
}

func (o OvaTaskAPI) RemoveTasksV1(ctx context.Context, in *desc.RemoveTaskV1Request) (*emptypb.Empty, error) {
	fmt.Println("RemoveTasksV1", in)
	return &emptypb.Empty{}, nil
}
