package auth_test

import (
	"testing"
	"time"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"github.com/0B1t322/service.auth/models/user"
	"github.com/0B1t322/service.auth/pkg/auth"
)

func TestFunc_GetJWT(t *testing.T) {
	db.Init()
	
	err := user.NewUser("dandemin", "1234", "admin").AddUser()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	u, err := user.GetUserByUserName("dandemin")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	authManaget := auth.NewAuthManager([]byte("secket_key"), "123", 1*time.Second)

	token, err := authManaget.CreateToken(u)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	info, err := auth.ParseToken(token, []byte("secket_key"))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if info == nil {
		t.Log("Info is nil")
		t.FailNow()
	}

	t.Log(info.GetUsername(), info.GetRole())

	timer := time.NewTimer(1*time.Second)
	<-timer.C

	info, err = auth.ParseToken(token, []byte("secket_key") )
	if err == nil {
		t.Fail()
	}

	u.DeleteUser()
}