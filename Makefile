gen_proto:
	protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_out=paths=source_relative:. --openapiv2_out . internal/proto/shtrafov-net-task/shtrafov-net-task.proto