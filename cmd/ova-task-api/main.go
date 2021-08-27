package main

import (
	"fmt"
	//ova_task_api "ozonva/ova-task-api/pkg/api/ova-task-api"
)

const configFilePath = "configs/ova-task-api.config"
const configUpdatePeriodSeconds = 1

func main() {
	//listen, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//server := grpc.NewServer()
	//ova_task_api.RegisterOvaTaskApiServer(server, api.NewOvaTaskApi())
	//
	//if err := server.Serve(listen); err != nil {
	//	fmt.Println("error", err)
	//}
}

func configUpdateHandle(configVersion string) {
	fmt.Printf("Config version: %v\r\n", configVersion)
}
