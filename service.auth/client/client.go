package client

import (
	"errors"
	"io"

	pkg_client "github.com/0B1t322/distanceLearningWebSite/pkg/client"

	pb "github.com/0B1t322/service.auth/authservice"
	"google.golang.org/grpc"
)

type Client interface {
	io.Closer
	pb.AuthServiceClient
}

func NewClient(
	network string,
	port string, 
	opts []grpc.DialOption,
) (Client, error) {
	c, conn, err := pkg_client.New(network, port, opts, pb.NewAuthServiceClient)
	if err != nil {
		return nil, err
	}

	client, ok := c.(pb.AuthServiceClient)
	if !ok {
		return nil, errors.New("Not okay")
	}

	return &struct{
		io.Closer
		pb.AuthServiceClient
	}{
		conn,
		client,
	}, nil
}
