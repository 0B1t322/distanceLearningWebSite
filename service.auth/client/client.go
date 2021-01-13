package client

import (
	"io"

	pb "github.com/0B1t322/service.auth/authservice"
	"google.golang.org/grpc"
)

type authServiceClient interface {
	pb.AuthServiceClient
}

type client struct {
	authServiceClient

	port 	string
	network string
	conn io.Closer
}

func NewClient(
	network string,
	port string, 
	opts []grpc.DialOption,
) (*client, error) {
	c := &client{network: network, port: port}
	conn, err := c.newConn(opts)
	if err != nil {
		return nil, err
	}
	
	c.conn = conn
	c.authServiceClient = pb.NewAuthServiceClient(conn)
	
	return c, nil
}

func (c *client) newConn(
	opts []grpc.DialOption,
) (*grpc.ClientConn, error) {
	return grpc.Dial(c.network+":"+c.port, opts...)
}

func (c *client) Close() error {
	return c.conn.Close()
}
