package client_test

import (
	"context"
	"testing"

	"github.com/0B1t322/distanceLearningWebSite/pkg/models/user"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"

	"google.golang.org/grpc/metadata"

	pb "github.com/0B1t322/service.auth/authservice"

	"github.com/0B1t322/service.auth/client"
	"google.golang.org/grpc"
)

func TestFunc_SignIn(t *testing.T) {
	db.Init()
	u := user.NewUser("dandem", "123", "admin")

	c, err := client.NewClient("127.0.0.1", "5050", []grpc.DialOption{ grpc.WithInsecure()})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	md := metadata.Pairs("role", u.Role)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := c.SignIn(
		ctx,
		&pb.AuthRequest{
			Username: u.Username,
			Password: u.Password,
		},
	)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Log(res)

	err = u.DeleteUser()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_SignUp(t *testing.T) {
	db.Init()
	u := user.NewUser("dandem", "123", "admin")

	c, err := client.NewClient("127.0.0.1", "5050", []grpc.DialOption{ grpc.WithInsecure()})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()


	md := metadata.Pairs("role", u.Role)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	_, err = c.SignIn(
		ctx,
		&pb.AuthRequest{
			Username: u.Username,
			Password: u.Password,
		},
	)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	res, err := c.SignUp(
		context.Background(),
		&pb.AuthRequest{
			Username: u.Username,
			Password: u.Password,
		},
	)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(res)

	err = u.DeleteUser()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

/*
func TestFunc_Interceptor(t *testing.T) {
	db.Init()

	u := user.NewUser("sss", "123", "user")
	c, err := client.NewClient("127.0.0.1", "5050", []grpc.DialOption{ grpc.WithInsecure()})
	if err != nil {
		t.Log(err)
		t.Log("Require to lauch server or push interceptor")
		t.SkipNow()
	}
	defer c.Close()

	md := metadata.Pairs("role", u.Role)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := c.SignIn(ctx, &pb.AuthRequest{Username: u.Username, Password: u.Password})

	if err == nil {
		t.Log(res)
		t.FailNow()
	}
	
}
*/

func TestFunc_Check(t *testing.T) {
	db.Init()

	u := user.NewUser("dandemin", "123", "user")
	c, err  := client.NewClient("127.0.0.1", "5050", []grpc.DialOption{ grpc.WithInsecure()})
	if err != nil {
		t.Log(err)
		t.Log("Require to lauch server or push interceptor")
		t.SkipNow()
	}
	defer c.Close()

	md := metadata.Pairs("role", u.Role)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := c.SignIn(ctx, &pb.AuthRequest{Username: u.Username, Password: u.Password})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	tokenInfo, err := c.Check(context.Background(), &pb.Token{Token: res.Token})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(tokenInfo)
	u.DeleteUser()
}