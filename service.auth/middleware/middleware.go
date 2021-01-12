package middleware

import (
	"strings"
	"google.golang.org/grpc/metadata"
	"github.com/0B1t322/auth-service/pkg/auth"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"context"
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

func CheckTokenInteceptor(
	ctx context.Context, 
	req interface{}, 
	info *grpc.UnaryServerInfo, 
	handler grpc.UnaryHandler,
	) (interface{}, error) {
		fnNotAuth := func(ctx context.Context) (interface{}, error) {
			ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("auth", "false"))
			return CheckAuth(ctx, req, info, handler)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return fnNotAuth(ctx)
		}

		var token string
		if t := md.Get("token"); len(t) > 0 {
			token = t[0]
		} else {
			return fnNotAuth(ctx)
		}
		var secretKey string
		if sk := md.Get("secretKey"); len(sk) > 0  {
			secretKey = sk[0]
		}

		tokenInfo, err := auth.ParseToken(token, []byte(secretKey))
		if err ==  auth.ErrInvalidToken {
			return fnNotAuth(ctx)
		} else if err != nil {
			log.Errorln("Error on checkInterceptor:",err)
			return fnNotAuth(ctx)
		}

		ctx = metadata.NewOutgoingContext(
			ctx, 
			metadata.Pairs(
				"auth", "true", 
				"role", tokenInfo.GetRole(), 
				"username", tokenInfo.GetUsername(),
			),
		)

		return CheckAuth(ctx, req, info, handler)
}

func CheckAuth(
	ctx context.Context, 
	req interface{}, 
	info *grpc.UnaryServerInfo, 
	handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			log.Warn("some going worng on parse  metadata")
		}

		var ifAuth string
		if a := md.Get("auth"); len(a) > 0 {
			ifAuth = a[0]
		}

		if strings.EqualFold(ifAuth, "false") {
			return nil, auth.ErrInvalidToken
		}

		return handler(ctx, req)
}