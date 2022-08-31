package grpcServer

import (
	"fmt"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"shtrafov-net-task/internal/infrastructure/proto/shtrafov-net-task"
	"shtrafov-net-task/internal/middleware"
)

func StartGRPCServer(port string, server grpcService.INNServiceServer) error {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		return fmt.Errorf("Error in listen %s", err.Error())
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			middleware.NewUnaryServerInterceptor(),
		)),
	)

	grpcService.RegisterINNServiceServer(grpcServer, server)

	log.Info().Msgf("Start grpc server...")

	return grpcServer.Serve(listen)
}
