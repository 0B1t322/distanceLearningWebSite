package main

import (
	"context"
	"strings"

	"flag"
	"fmt"
	"net"

	pb "github.com/0B1t322/distanceLearningWebSite/protos/coursesservice"
	"github.com/0B1t322/service.courses/server"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"github.com/0B1t322/distanceLearningWebSite/pkg/middleware"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	port 		= 	flag.String("port", "5050", "start grpc server on this port")
	secretKey 	= 	flag.String("sk", "key","use for check user token")
	logger		=	log.New()
)

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		logger.Panicln(err)
	}

	opts := []grpc.ServerOption {
		grpc.ChainUnaryInterceptor(
			middleware.TokenParsesInterceptor(*secretKey),
			middleware.ErrorLoggerUnaryInterceptor(logger),
			// TODO make internal pkg for this bulding interceptrit
			middleware.MethodsCheckerUnaryInterceptor(
				func(ctx context.Context) bool {
					md, ok := metadata.FromIncomingContext(ctx)
					if !ok {
						logger.Warn("Error on methodchecker: cant take metadata")
						return false
					}
					role := md.Get("role")[0]
					logger.Info(role)
					if !strings.EqualFold(role, "admin") {
						return false
					}

					return  true
				},
				func() map[string]struct{} {
					m := make(map[string]struct{})
					m["addcourse"] = struct{}{}

					return m
				}(),
			),
		),
	}

	DB, err := db.DBManger.OpenDataBase("courses")
	if err != nil {
		logger.Panicln(err)
	}

	grpcServer := grpc.NewServer(opts...)

	pb.RegisterCoursesServiceServer(
		grpcServer, 
		server.NewServer(
			DB,
		),
	)

	logger.Infof("Starting grpc server on: :%s\n", *port)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Panicf("failed to start server: %v\n",err)
	}
	
}