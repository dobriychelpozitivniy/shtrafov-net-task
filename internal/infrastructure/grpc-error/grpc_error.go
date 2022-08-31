package grpcError

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func New(code codes.Code, msg string, err error) error {
	if err != nil {
		msg = fmt.Sprintf(msg+": %s", err.Error())
	}

	return status.New(code, msg).Err()
}
