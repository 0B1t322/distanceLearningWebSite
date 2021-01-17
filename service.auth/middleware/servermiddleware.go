package middleware

import (
	"context"

	pb "github.com/0B1t322/service.auth/authservice"
	"github.com/0B1t322/service.auth/client"
	"github.com/0B1t322/service.auth/pkg/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
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


/* 
CheckAuthUnaryServerInterceptor return a UnaryServerInterceptor
	require a client interceptor: TokenUnaryInterceptor
	token in ctx with key "token"
	params:
		netwotk - adress to grpc auth server
		port 	- port to grpc auth server
		opts	- some dial opts that you need
*/
func CheckAuthUnaryServerInterceptor(
	network string,
	port  	string,
	opts ...grpc.DialOption,
) (grpc.UnaryServerInterceptor) {
	interceptor := func(
		ctx context.Context, 
		req interface{}, 
		info *grpc.UnaryServerInfo, 
		handler grpc.UnaryHandler,
	)  (interface{}, error) {
			c, err := client.NewClient(network, port, opts)
			if err != nil {
				log.Warn(err)
				return nil, grpc.Errorf(codes.Internal, "interal server error")
			}
			defer c.Close()

			var token string
			md, ok := metadata.FromIncomingContext(ctx)
			if !ok {
				log.Warn("No token in ctx")
				token = ""
			}

			if t := md.Get("token"); len(t) > 0 {
				token = t[0]
			} else {
				log.Warn("can't take token from metadata")
				token = ""
			}
		
			_, err = c.Check(ctx, &pb.Token{Token: token})
			if err == auth.ErrInvalidToken {
				return nil, grpc.Errorf(codes.Unauthenticated, "%v", err)
			} else if err == auth.ErrTokenExpire {
				return nil, grpc.Errorf(codes.Unauthenticated, "%v", err)
			} else if err != nil {
				log.Warn(err)
				return nil, grpc.Errorf(codes.Internal, "interal server error")
			}
			return handler(ctx, req)
		
		}
	

	return interceptor
}