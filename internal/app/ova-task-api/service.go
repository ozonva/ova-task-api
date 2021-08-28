package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	desc "ozonva/ova-task-api/pkg/api/ova-task-api"
)

type OvaTaskAPI struct {
	desc.UnimplementedOvaTaskApiServer
}

func NewOvaTaskApi() desc.OvaTaskApiServer {
	return &OvaTaskAPI{}
}

func (o OvaTaskAPI) Ping(ctx context.Context, in *emptypb.Empty) (*desc.StringMessage, error) {
	return &desc.StringMessage{Value: "Hello from API"}, nil
}

func (o OvaTaskAPI) Echo(ctx context.Context, in *desc.StringMessage) (*desc.StringMessage, error) {
	return in, nil
}
