package client_test

import (
	"context"
	"testing"

	grpc_middleware "github.com/0B1t322/distanceLearningWebSite/pkg/middleware"
	"github.com/0B1t322/distanceLearningWebSite/service.auth/authservice"
	auth_client "github.com/0B1t322/distanceLearningWebSite/service.auth/client"
	"github.com/0B1t322/service.courses/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func init() {
	c, err := auth_client.NewClient("127.0.0.1","5050", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err)
	}
	defer c.Close()
	
	authC, _ := c.(authservice.AuthServiceClient)
	
	ctx :=  metadata.NewOutgoingContext(context.Background(), metadata.Pairs("role", "admin"))

	resp, err := authC.SignUp(ctx, &authservice.AuthRequest{Username: "dan", Password: "123"})
	if err != nil {
		panic(err)
	}

	opts = append(opts, grpc.WithInsecure(), grpc.WithUnaryInterceptor(grpc_middleware.TokenUnaryClientInterceptor(resp.Token)))
}

var opts []grpc.DialOption


func TestFunc_NewClient(t *testing.T) {
	c, err :=client.NewClient("127.0.0.1","5051", opts)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	
	defer c.Close()
	
}

func TestFunc_AddCourse(t *testing.T) {

}