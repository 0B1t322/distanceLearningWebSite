package user_test

import (
	"time"
	"fmt"
	"testing"

	uc "github.com/0B1t322/distanceLearningWebSite/pkg/controllers/user"
	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
	um "github.com/0B1t322/distanceLearningWebSite/pkg/models/user"
)

var (
	DBManger = db.NewManager(
		"root",
		"root",
		"127.0.0.1:3306",
		15 * time.Second,
	)
)


func TestFUNC_AddUser(t *testing.T) {
	DB, _ := DBManger.OpenDataBase("auth")
	c := uc.NewUserController(DB)

	err := c.AddUser(um.NewUser("dandemin", "123", "user"))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = c.AddUser(um.NewUser("dandemin", "123", "user"))
	if err != uc.ErrUserExsist {
		t.Log(err)
		t.FailNow()
	}

	u := um.NewUser("dandemin", "123", "user")
	err = c.DeleteUser(u)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestFUNC_GetUserByID(t *testing.T) {
	DB, _ := DBManger.OpenDataBase("auth")
	c := uc.NewUserController(DB)

	_, err := c.GetUserByID("-1")

	if err != uc.ErrUserNotFound {
	 	t.FailNow()
	}

	err = c.AddUser(um.NewUser("dandemin", "123", "admin"))
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	_u, err := c.GetUserByUserName("dandemin")
	if  err  != nil {
		t.Log(err)
		t.FailNow()
	}

	u, err := c.GetUserByID(fmt.Sprint(_u.ID))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log(u)
	c.DeleteUser(u)
}

func TestFUNC_GetUserByUserName(t *testing.T) {
	DB, _ := DBManger.OpenDataBase("auth")
	c := uc.NewUserController(DB)

	_, err := c.GetUserByUserName("dandemin")
	if err != uc.ErrUserNotFound {
		t.Log(err)
		t.FailNow()
	}

	err = c.AddUser(um.NewUser("dandemin", "123", "admin"))
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	u, err := c.GetUserByUserName("dandemin")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	err = c.DeleteUser(u)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}

func TestFunc_DeleteUser(t *testing.T) {
	DB, _ := DBManger.OpenDataBase("auth")
	c := uc.NewUserController(DB)

	u := um.NewUser("dandemin", "123", "admin")

	err := c.AddUser(u)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	u, err = c.GetUserByUserName("dandemin")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	t.Log(u)

	err = c.DeleteUser(u)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	u, err = c.GetUserByUserName("dandemin")
	t.Log(u)
}

