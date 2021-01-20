package marhsall_test

import (
	"reflect"
	"github.com/0B1t322/distanceLearningWebSite/pkg/marshall"
	"testing"
)

func TestFunc_Marshall(t *testing.T) {
	type f struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}

	type a struct {
		Name string `json:"name"`
	}

	F := &f{Name: "cool name", Role: "God"}
	A := &a{}
	err := marshall.Marshall(F, A)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	t.Log("F: ", F)
	t.Log("A: ", A)
}

func TestFunc_Reflection(t *testing.T) {
	type user struct {
		Name 		string	`json:"name"`
		Password	string	`json:"password"`
		Role 		string 	`json:"role"`
	}

	type pbUser struct {
		Name	string `json:"name"`
		Role	string `json:"role"`	
	}

	u := &user{Name: "some_name", Password: "some_pass", Role: "some_role"}

	st := reflect.TypeOf(u).Elem()
	fields := st.NumField()

	for i := 0; i < fields; i++ {
		t.Log(st.Field(i).Tag)
	}
}

// func marsh(from interface{}, to interface{}) {
// 	fromStruct 	:= reflect.TypeOf(from).Elem()
// 	toStruct 	:= reflect.TypeOf(to).Elem()
// }