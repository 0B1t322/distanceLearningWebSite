package middleware

import (
	"github.com/0B1t322/distanceLearningWebSite/pkg/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/metadata"
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

/*
TokenParsesInterceptor parse token and put token info into metadata
*/
func TokenParsesInterceptor(sk string) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context, 
		req interface{}, 
		info *grpc.UnaryServerInfo, 
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		var token string
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "don't get token from client")
		}

		if t := md.Get("token"); len(t) > 0 {
			token = t[0]
		} else {
			return nil, status.Error(codes.Unauthenticated, "don't get token from client")
		}

		tokenInfo, err := auth.ParseToken(token, []byte(sk))
		if err == auth.ErrInvalidToken {
			return nil, grpc.Errorf(codes.Unauthenticated, "%v", err)
		} else if err == auth.ErrTokenExpire {
			return nil, grpc.Errorf(codes.Unauthenticated, "%v", err)
		} else if err != nil {
			return nil, grpc.Errorf(codes.Internal, "%v", err)
		}
		

		ctx = metadata.NewIncomingContext(
			ctx, 
			metadata.Pairs(
				"uid", 		tokenInfo.UID,
				"username", tokenInfo.Username,
				"role", 	tokenInfo.Role,	
			),
		)

		return handler(ctx, req)
	}
}