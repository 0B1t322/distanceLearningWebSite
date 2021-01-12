package main

import (
	"time"
	"github.com/0B1t322/auth-service/pkg/auth"
	"github.com/0B1t322/auth-service/server"
	pb "github.com/0B1t322/auth-service/authservice"
	"flag"
	"fmt"
	"net"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port 		= flag.String("port", ":8080", "start grpc server on this port")
	secretKey	= flag.String("sk", "my_secret_key", "secret key - need to hash JWT token")
)

func main() {
	// db.Init() alredy parse flags so we don't need to write this again
	db.Init()
	log.Info("db init all is okay!")

	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Panicln(err)
	}

	opts := []grpc.ServerOption{}

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterAuthServiceServer(
		grpcServer, 
		server.NewServer(
			auth.NewAuthManager(
				[]byte(*secretKey),
				"some_salt",
				2 * time.Hour,
			),
		),
	)

	log.Infof("Starting grpc server on: %s\n", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Panicf("failed to start server: %v",err)
	}

}

// TODO добавить работу с бд tokens
// TODO удаление пользователей