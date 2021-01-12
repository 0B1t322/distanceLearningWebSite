package user_test

import (
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"github.com/0B1t322/auth-service/models/user"
	
	"fmt"
	"testing"
)

func BenchMarkDeleteUser(b *testing.B) {
	db.Init()
	for i := 0; i < b.N; i++ {
		u := user.NewUser(
			fmt.Sprintf("dandemin%v",i),
			"123",
			"user",
		)

		err := u.AddUser()
		if err != nil {
			b.Log(err)
			b.FailNow()
		}

		defer u.DeleteUser()
	}
	b.ResetTimer()
}