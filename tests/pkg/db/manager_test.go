package db_test

import (
	"testing"
	"time"

	"github.com/0B1t322/distanceLearningWebSite/pkg/db"
)

func TestFunc_NewManager(t *testing.T) {
	// const pathToJSON = "list.json"
	// file, err := os.Open(pathToJSON)
	// if err != nil {
	// 	t.Log(err)
	// 	t.FailNow()
	// }
	
	manager := db.NewManager("user", "pass","network", time.Second)

	t.Log(manager)
}

func TestFunc_OpenDataBase(t *testing.T) {
	manager := db.NewManager("root", "root","127.0.0.1:3306", time.Second)

	_, err := manager.OpenDataBase("auth")
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	
}