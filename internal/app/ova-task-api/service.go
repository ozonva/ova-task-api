package api

import (
	"context"
	"encoding/json"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("CreateTaskV1")
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (o OvaTaskAPI) DescribeTaskV1(ctx context.Context, in *desc.DescribeTaskV1Request) (*desc.DescribeTaskV1Response, error) {
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("DescribeTaskV1")
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (o OvaTaskAPI) ListTasksV1(ctx context.Context, in *desc.ListTasksV1Request) (*desc.ListTasksV1Response, error) {
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("ListTasksV1")
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (o OvaTaskAPI) RemoveTasksV1(ctx context.Context, in *desc.RemoveTaskV1Request) (*emptypb.Empty, error) {
	inJson, _ := json.Marshal(in)
	log.Debug().RawJSON("in", inJson).Msg("RemoveTasksV1")
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
