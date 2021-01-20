package client_test

import (
	"context"
	"testing"

	"github.com/0B1t322/service.auth/pkg/auth"
	"github.com/0B1t322/service.auth/server"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"github.com/0B1t322/distanceLearningWebSite/pkg/models/user"

	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/0B1t322/service.auth/authservice"

	"github.com/0B1t322/service.auth/client"
	"google.golang.org/grpc"
)

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

	res, err := c.SignUp(
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

func  TestFunc_SingUp_UserExsist(t *testing.T) {
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

	_, err = c.SignUp(
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

	res, err := c.SignUp(
		ctx, 
		&pb.AuthRequest{
			Username: u.Username,
			Password: u.Password,
		},
	)

	if err == nil {
		t.Log("Faield with user exsist")
		t.FailNow()
	}
	s, ok := status.FromError(err)
	if ok {
		t.Log("code: ",s.Code())
		t.Log("desc: ",s.Message())
	}
	t.Log("err: ",err)
	t.Log("res: ",res)
	

	err = u.DeleteUser()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_SignUp_ErrGetRoleFromCTX(t *testing.T) {
	db.Init()

	u := user.NewUser("dandem", "123", "admin")

	c, err := client.NewClient("127.0.0.1", "5050", []grpc.DialOption{ grpc.WithInsecure()})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer c.Close()

	_, err = c.SignUp(
		context.Background(),
		&pb.AuthRequest{
			Username: u.Username,
			Password: u.Password,
		},
	)

	if err == nil {
		t.Log("Don't get err")
		t.FailNow()
	}

	t.Log(err)
}

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
	_, err = c.SignUp(
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

	res, err := c.SignIn(
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

func TestFunc_SingIn_Unauthenticated(t *testing.T) {
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

	_, err = c.SignUp(
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

	_, err = c.SignIn(
		context.Background(),
		&pb.AuthRequest{
			Username: "ddd",
			Password: "123",
		},
	)
	if err == nil {
		t.Log("Dont get   errros")
		t.FailNow()
	}

	s, ok := status.FromError(err)
	if !ok {
		t.Log("Cant  get status")
		t.FailNow()
	}
	if s.Message() != server.ErrIncorrectUserNamePass.Error() {
		t.Log("Unexcpectet error")
		t.Log(err)
		t.FailNow()
	}

	_, err = c.SignIn(
		context.Background(),
		&pb.AuthRequest{
			Username: "dandem",
			Password: "1",
		},
	)
	if err == nil {
		t.Log("Dont get   errros")
		t.FailNow()
	}

	s, ok = status.FromError(err)
	if !ok {
		t.Log("Cant  get status")
		t.FailNow()
	}
	if s.Message() != server.ErrIncorrectUserNamePass.Error() {
		t.Log("Unexcpectet error")
		t.Log(err)
		t.FailNow()
	}

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

	res, err := c.SignUp(ctx, &pb.AuthRequest{Username: u.Username, Password: u.Password})
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

func TestFunc_Check_InvalidArgument(t *testing.T) {
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

	res, err := c.SignUp(ctx, &pb.AuthRequest{Username: u.Username, Password: u.Password})
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	_, err = c.Check(context.Background(), &pb.Token{Token: res.Token + "2"})
	if err == nil {
		t.Log("Dont get error")
		t.FailNow()
	}

	s, ok := status.FromError(err)
	if !ok {
		t.Log("Cant  get status")
		t.FailNow()
	}
	if s.Message() != auth.ErrInvalidToken.Error() {
		t.Log("Unexcpected error")
		t.Log(err)
		t.FailNow()
	}

	t.Log(err)
	u.DeleteUser()
}