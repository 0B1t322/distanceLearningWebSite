package client_test

import (
	"time"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"testing"
	pb "github.com/0B1t322/service.auth/authservice"
	"google.golang.org/grpc"
	"github.com/0B1t322/service.auth/client"
	"context"
	"github.com/0B1t322/distanceLearningWebSite/pkg/models/user"
	uc "github.com/0B1t322/distanceLearningWebSite/pkg/controllers/user"
	"fmt"
	"google.golang.org/grpc/metadata"
)

var (
	DBManger = db.NewManager(
		"root",
		"root",
		"127.0.0.1:3306",
		15 * time.Second,
	)
)


func BenchSignUp(b *testing.B) {
	DB, _ := DBManger.OpenDataBase("auth")
	controll := uc.NewUserController(DB)
	for i := 0; i < b.N; i++ {
		u := user.NewUser(fmt.Sprintf("dandemin%v", i),"1","user")

		c, err := client.NewClient("127.0.0.1", "5050", []grpc.DialOption{ grpc.WithInsecure()})
		if err != nil {
			b.Log(err)
			b.FailNow()
		}
		defer c.Close()

		md := metadata.Pairs("rol", u.Role)
		ctx := metadata.NewOutgoingContext(context.Background(), md)

		res, err := c.SignUp(
			ctx,
			&pb.AuthRequest{
				Username: u.Username,
				Password: u.Password,
			},
		)

		if err != nil {
			b.Log(err)
			b.Fail()
		}
		fmt.Print("Hello")
		b.Log(res)

		defer func(b *testing.B) {
			err = controll.DeleteUser(u)
			if err != nil {
				b.Log(err)
				b.FailNow()
			}
		}(b)
	}
}