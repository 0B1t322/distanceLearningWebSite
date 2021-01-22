package auth_test

import (
	uc "github.com/0B1t322/distanceLearningWebSite/pkg/controllers/user"
	"errors"
	"testing"
	"time"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"github.com/0B1t322/distanceLearningWebSite/pkg/models/user"
	"github.com/0B1t322/service.auth/pkg/auth"
)

var (
	DBManger = db.NewManager(
		"root",
		"root",
		"127.0.0.1:3306",
		15 * time.Second,
	)
)

func TestFunc_GetJWT(t *testing.T) {
	DB, _ := DBManger.OpenDataBase("auth")
	controll := uc.NewUserController(DB)
	
	err := controll.AddUser(user.NewUser("dandemin", "1234", "admin"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	u, err := controll.GetUserByUserName("dandemin")
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

	t.Log(info.Username, info.Role, info.UID)

	timer := time.NewTimer(1*time.Second)
	<-timer.C

	info, err = auth.ParseToken(token, []byte("secket_key") )
	if err == nil {
		t.Fail()
	}
	t.Log(err)
	t.Log(errors.Unwrap(err))

	controll.DeleteUser(u)
}