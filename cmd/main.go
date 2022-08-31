package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"shtrafov-net-task/internal/config"
	grpcHttpServer "shtrafov-net-task/internal/infrastructure/grpc-http-server"
	grpcServer "shtrafov-net-task/internal/infrastructure/grpc-server"
	"shtrafov-net-task/internal/rusprofile-service"
)

func main() {
	cfgPath := flag.String("config", "configs/config.toml", "path to config file")
	flag.Parse()

	cfg, err := config.New(*cfgPath)
	if err != nil {
		log.Panic().Msgf("Error create config: %s", err)
	}

	service := rusprofileService.New(cfg.RusporfileService.Url)
	go func() {
		err := grpcServer.StartGRPCServer(cfg.App.GrpcServerPort, service)
		if err != nil {
			log.Panic().Msgf("Error when start grpc server: %s", err)
		}
	}()

	err = grpcHttpServer.StartGrpcHttpServer(cfg.App.GrpcServerPort, cfg.App.GrpcHttpServerPort)
	if err != nil {
		log.Panic().Msgf("Error when start grpc http server: %s", err)
	}
}
