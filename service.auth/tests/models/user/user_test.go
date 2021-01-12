package user_test

import (
	"fmt"
	"testing"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	"github.com/0B1t322/auth-service/models/user"
)

func TestFUNC_AddUser(t *testing.T) {
	db.Init()
	err := user.NewUser("dandemin", "123", "user").AddUser()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	err = user.NewUser("dandemin", "123", "user").AddUser()
	if err != user.ErrUserExsist {
		t.Log(err)
		t.Fail()
	}

	u := user.NewUser("dandemin", "123", "user")
	err = u.DeleteUser()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestFUNC_GetUserByID(t *testing.T) {
	db.Init()
	_, err := user.GetUserByID("-1")

	if err != user.ErrUserNotFound {
	 	t.Fail()
	}

	err = user.NewUser("dandemin", "123", "admin").AddUser()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	_u, err := user.GetUserByUserName("dandemin")
	if  err  != nil {
		t.Log(err)
		t.FailNow()
	}

	u, err := user.GetUserByID(fmt.Sprint(_u.ID))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(u)
	u.DeleteUser()
}

func TestFUNC_GetUserByUserName(t *testing.T) {
	db.Init()

	_, err := user.GetUserByUserName("dandemin")
	if err != user.ErrUserNotFound {
		t.Log(err)
		t.FailNow()
	}

	err = user.NewUser("dandemin", "123", "admin").AddUser()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	u, err := user.GetUserByUserName("dandemin")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = u.DeleteUser()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_deleteUser(t *testing.T) {
	db.Init()
	u := user.NewUser("dandemin", "123", "admin")

	err := u.AddUser()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	u, err = user.GetUserByUserName("dandemin")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(u)

	err = u.DeleteUser()
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	u, err = user.GetUserByUserName("dandemin")
	t.Log(u)
}

//TODO рефактор тестов