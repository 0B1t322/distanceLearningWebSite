package main

import (
	"github.com/0B1t322/distanceLearningWebSite/pkg/auth"
	"github.com/0B1t322/distanceLearningWebSite/pkg/middleware"
	"time"

	"github.com/0B1t322/service.auth/server"
	pb "github.com/0B1t322/distanceLearningWebSite/protos/authservice"
	"flag"
	"fmt"
	"net"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port 		= 	flag.String("port", "5050", "start grpc server on this port")
	secretKey	= 	flag.String("sk", "my_secret_key", "secret key - need to hash JWT token")
	logger 		=	log.New()
)


func main() {
	flag.Parse()

	logger.ReportCaller = true
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		logger.Panicln(err)
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			middleware.ErrorLoggerUnaryInterceptor(logger),
		),
	}

	DB, err := db.DBManger.OpenDataBase("auth")
	if err != nil {
		logger.Panic(err)
	}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAuthServiceServer(
		grpcServer, 
		server.NewServer(
			auth.NewAuthManager(
				[]byte(*secretKey),
				"some_salt",
				2 * time.Hour,
			),
			DB,
		),
	)

	logger.Infof("Starting grpc server on: %s\n", *port)

	if err := grpcServer.Serve(lis); err != nil {
		logger.Panicf("failed to start server: %v",err)
	}

	

}

// TODO добавить работу с бд tokens
// TODO удаление пользователей