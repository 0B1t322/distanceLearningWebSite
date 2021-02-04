package middleware

import (
	"context"
	"google.golang.org/grpc"
	log "github.com/sirupsen/logrus"
)

/*
ErrorLoggerUnaryInterceptor log if cause some error on the handler

If using a chain should be last
*/
func ErrorLoggerUnaryInterceptor(l *log.Logger) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, 
		req interface{}, 
		info *grpc.UnaryServerInfo, 
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			l.Errorf("%v", err)
		}

		return resp, err
	}
}

/*
do not use not done
*/
func ErrLoggerStreamInterceptor(l *log.Logger) grpc.StreamServerInterceptor {
	return func(
		srv interface{}, 
		ss grpc.ServerStream, 
		info *grpc.StreamServerInfo, 
		handler grpc.StreamHandler,
	) error {
		// TODO 
		return nil
	}
}