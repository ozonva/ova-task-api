package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"ozonva/ova-task-api/internal/app/ova-task-api"
	"ozonva/ova-task-api/internal/utils"
	apiServer "ozonva/ova-task-api/pkg/api/ova-task-api"
	"strconv"
)

const configFilePath = "configs/ova-task-api.config"
const configUpdatePeriodSeconds = 1

func main() {
	config, err := utils.ReadConfig(configFilePath)
	if err != nil {
		panic(err)
	}
	go runJSON(config.Grpc.Port, config.Http.Port)
	if err := run(config.Grpc.Port); err != nil {
		log.Fatal(err)
	}
}

func runJSON(grpcPort int, httpPort int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	grpcAddress := "localhost:" + strconv.Itoa(grpcPort)
	err := apiServer.RegisterOvaTaskApiHandlerFromEndpoint(ctx, mux, grpcAddress, opts)
	if err != nil {
		panic(err)
	}

	err = http.ListenAndServe(":"+strconv.Itoa(httpPort), mux)
	if err != nil {
		panic(err)
	}
}

func run(grpcPort int) error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server is listening...")
	server := grpc.NewServer()
	apiServer.RegisterOvaTaskApiServer(server, api.NewOvaTaskApi())

	if err := server.Serve(listener); err != nil {
		fmt.Println("error", err)
	}
	return nil
}

func configUpdateHandle(configVersion string) {
	fmt.Printf("Config version: %v\r\n", configVersion)
}
