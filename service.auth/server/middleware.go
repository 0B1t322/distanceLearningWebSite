package server

import (
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