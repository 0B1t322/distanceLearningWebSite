package middleware

import (
	"google.golang.org/grpc"
	"context"

	"google.golang.org/grpc/metadata"
)

/* 
TokenUnaryClientInterceptor return UnaryClientInterceptor
	put token into context with key "token" with grpc metadata
	params:
		token - access JWT token
*/
func TokenUnaryClientInterceptor(token string) (grpc.UnaryClientInterceptor) {
	interceptor := func (
		ctx context.Context,
		method string,
		req interface{},
		reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("token", token))
		err := invoker(ctx, method, req, reply, cc, opts...)

		return err
	}

	return interceptor
}