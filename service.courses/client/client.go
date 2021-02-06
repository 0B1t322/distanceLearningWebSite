package client

import (
	"errors"
	"io"

	pkg_client "github.com/0B1t322/distanceLearningWebSite/pkg/client"

	pb "github.com/0B1t322/service.courses/coursesservice"
	"google.golang.org/grpc"
)

type Client interface {
	pb.CoursesServiceClient
	io.Closer
}

func NewClient(
	network string,
	port string, 
	opts []grpc.DialOption,
) (Client, error) {
	c, conn, err := pkg_client.New(network, port, opts,pb.NewCoursesServiceClient)
	if err != nil {
		return nil, err
	}

	client, ok := c.(pb.CoursesServiceClient)
	if !ok {
		return nil, errors.New("Not okay")
	}

	return &struct{
		pb.CoursesServiceClient
		io.Closer
	}{
		client,
		conn,
	}, nil
}