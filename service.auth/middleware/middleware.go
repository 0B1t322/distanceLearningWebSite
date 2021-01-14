package middleware

import (
	"context"

	pb "github.com/0B1t322/service.auth/authservice"
	"github.com/0B1t322/service.auth/client"
	"github.com/0B1t322/service.auth/pkg/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)


func ErrorHandlerInteceptor(
	ctx context.Context, 
	req interface{}, 
	info *grpc.UnaryServerInfo, 
	handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Info("Interceptor")
		responce, err := handler(ctx, req)
		if err != nil {
			log.Errorln(err)			
		}
		
		return responce, err
}

func CheckAuthInterceptor(
	network string,
	port  	string,
	token	string,
	opts ...grpc.DialOption,
) (grpc.UnaryServerInterceptor) {
	interceptor := func(ctx context.Context, 
		req interface{}, 
		info *grpc.UnaryServerInfo, 
		handler grpc.UnaryHandler,
	)  (interface{}, error) {
		c, err := client.NewClient(network, port, opts)
		if err != nil {
			log.Warn(err)
			return nil, grpc.Errorf(codes.Internal, "interal server err")
		}
		defer c.Close()

		_, err = c.Check(ctx, &pb.Token{Token: token})
		if err  == auth.ErrInvalidToken {
			return nil, grpc.Errorf(codes.Unauthenticated, "Invalid token")
		}
		return handler(ctx, req)
	}
	// TODO check expire token

	return interceptor
}