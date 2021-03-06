package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"ozonva/ova-task-api/internal/app/ova-task-api"
	repopkg "ozonva/ova-task-api/internal/repo"
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

	dbConnectionString := fmt.Sprintf(
		"port=%v user=%v password=%v dbname=%v sslmode=disable",
		config.Db.Port,
		config.Db.User,
		config.Db.Password,
		config.Db.DataBase,
	)
	db, err := repopkg.OpenDb(dbConnectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load driver")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal().Msg("failed to close db")
		}
	}(db)

	repo := repopkg.NewRepo(db)
	go runHttpGateway(config.Grpc.Port, config.Http.Port)
	if err := runGrpc(config.Grpc.Port, repo); err != nil {
		log.Fatal().Err(err).Send()
	}
}

func runHttpGateway(grpcPort int, httpPort int) {
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

func runGrpc(grpcPort int, repo repopkg.Repo) error {
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(grpcPort))
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to listen")
	}
	log.Info().Msg("Server is listening...")
	server := grpc.NewServer()
	apiServer.RegisterOvaTaskApiServer(server, api.NewOvaTaskApi(repo))

	if err := server.Serve(listener); err != nil {
		log.Fatal().Err(err).Send()
	}
	return nil
}

func configUpdateHandle(configVersion string) {
	fmt.Printf("Config version: %v\r\n", configVersion)
}
