package client

import (
	"io"
	"reflect"

	"google.golang.org/grpc"
)

type Client struct {
	port 	string
	network	string
}

func New(
	network, 
	port string,
	opts []grpc.DialOption,
	NewServiceClient interface{},
) 	(interface{}, io.Closer ,error) {
	c := Client{network: network, port: port}
	conn, err := c.newConn(opts)
	if err != nil {
		return nil, nil, err
	}
	
	clientInterface := getClient(NewServiceClient, conn)


	return clientInterface, conn, nil
}

func (c *Client) newConn(
	opts []grpc.DialOption,
) (*grpc.ClientConn, error) {
	return grpc.Dial(c.adress(), opts...)
}

func (c *Client) adress() string {
	return c.network+":"+c.port
}

func getClient(NewServiceClient interface{}, conn grpc.ClientConnInterface) interface{} {
	return reflect.ValueOf(NewServiceClient).Call(
		[]reflect.Value {
			reflect.ValueOf(conn),
		},
	)[0].Interface()
}