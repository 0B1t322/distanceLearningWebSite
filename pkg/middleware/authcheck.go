package middleware

import (
	"context"

	"google.golang.org/grpc"
)

func AuthCheckInteceptor(
	ctx 		context.Context,
	req 		interface{},
	info 		*grpc.UnaryServerInfo,
	handler 	grpc.UnaryHandler,
) (interface{}, error) {
	// TODO
	return nil, nil
}
