package ova_task_api

import (
	desc "ozonva/ova-task-api/pkg/api/ova-task-api"
)

type OvaTaskAPI struct {
	desc.UnimplementedOvaTaskApiServer
}

func NewOvaTaskApi() desc.OvaTaskApiServer {
	return &OvaTaskAPI{}
}
