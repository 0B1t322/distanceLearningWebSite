package main

import (
	"time"
	"github.com/0B1t322/service.auth/pkg/auth"
	"github.com/0B1t322/service.auth/server"
	pb "github.com/0B1t322/service.auth/authservice"
	"flag"
	"fmt"
	"net"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	port 		= flag.String("port", "5050", "start grpc server on this port")
	secretKey	= flag.String("sk", "my_secret_key", "secret key - need to hash JWT token")
)


func main() {
	// db.Init() alredy parse flags so we don't need to write this again
	flag.Parse()
	
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Panicln(err)
	}

	opts := []grpc.ServerOption{
		
	}

	DB, err := db.DBManger.OpenDataBase("auth")
	if err != nil {
		log.Panic(err)
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

	log.Infof("Starting grpc server on: %s\n", *port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Panicf("failed to start server: %v",err)
	}

	

}

// TODO добавить работу с бд tokens
// TODO удаление пользователей