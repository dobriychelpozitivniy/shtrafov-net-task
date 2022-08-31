package grpcHttpServer

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
	"shtrafov-net-task/internal/infrastructure/proto/shtrafov-net-task"
)

func StartGrpcHttpServer(grpcServerEndpointPort, grpcHttpServerPort string) error {
	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := grpcService.RegisterINNServiceHandlerFromEndpoint(ctx, mux, grpcServerEndpointPort, opts)
	if err != nil {
		return err
	}

	log.Info().Msgf("Start grpc http server...")

	return http.ListenAndServe(grpcHttpServerPort, mux)
}
