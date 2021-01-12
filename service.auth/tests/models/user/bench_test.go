package user_test

import (
	"github.com/0B1t322/auth-service/models/user"
	"github.com/0B1t322/auth-service/db"
	"fmt"
	"testing"
)

func BenchMarkDeleteUser(b *testing.B) {
	db.Init(true)
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