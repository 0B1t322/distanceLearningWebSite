package marhsall_test

import (
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