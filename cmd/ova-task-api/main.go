package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	api "ozonva/ova-task-api/internal/app/ova-task-api"
	ova_task_api "ozonva/ova-task-api/pkg/api/ova-task-api"
)

const configFilePath = "configs/ova-task-api.config"
const configUpdatePeriodSeconds = 1

const (
	grpcPort           = ":82"
	grpcServerEndpoint = "localhost:82"
	httpPort           = ":8081"
)

func main() {
	go runJSON()

	if err := run(); err != nil {
		log.Fatal(err)
	}
}
func runJSON() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err := ova_task_api.RegisterOvaTaskApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(httpPort, mux)
	if err != nil {
		panic(err)
	}
}

func run() error {
	listener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server is listening...")
	server := grpc.NewServer()
	ova_task_api.RegisterOvaTaskApiServer(server, api.NewOvaTaskApi())

	if err := server.Serve(listener); err != nil {
		fmt.Println("error", err)
	}
	return nil
}

func configUpdateHandle(configVersion string) {
	fmt.Printf("Config version: %v\r\n", configVersion)
}
