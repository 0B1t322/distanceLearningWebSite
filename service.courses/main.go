package main

import (

	"github.com/0B1t322/service.courses/server"
	pb "github.com/0B1t322/distanceLearningWebSite/protos/coursesservice"
	"flag"
	"fmt"
	"net"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"github.com/0B1t322/distanceLearningWebSite/pkg/middleware"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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