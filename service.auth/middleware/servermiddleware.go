package middleware

import (
	"context"

	pb "github.com/0B1t322/service.auth/authservice"
	"github.com/0B1t322/service.auth/client"
	"github.com/0B1t322/service.auth/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

/*
CheckAuthUnaryServerInterceptor return a UnaryServerInterceptor
	require a client interceptor: TokenUnaryInterceptor
	token in ctx with key "token"
	launch handler with metadata with keys: "uid", "username", "role"
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
				return nil, grpc.Errorf(codes.Internal, "%v",err)
			}
			defer c.Close()

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
		
			tokenInfo, err := c.Check(ctx, &pb.Token{Token: token})
			if err == auth.ErrInvalidToken {
				return nil, grpc.Errorf(codes.Unauthenticated, "%v", err)
			} else if err == auth.ErrTokenExpire {
				return nil, grpc.Errorf(codes.Unauthenticated, "%v", err)
			} else if err != nil {
				return nil, grpc.Errorf(codes.Internal, "interal server error")
			}

			ctx = metadata.NewIncomingContext(
				ctx, 
				metadata.Pairs(
					"uid", 		tokenInfo.Uid,
					"username", tokenInfo.Username,
					"role", 	tokenInfo.Role,	
				),
			)

			return handler(ctx, req)
		
		}
	

	return interceptor
}

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