package marhsall_test

import (
	"fmt"
	"github.com/0B1t322/distanceLearningWebSite/pkg/marshall"
	"testing"
)

func BenchMark(b *testing.B) {
	type a struct {
		Name string `json:"name"`
	}
	type s struct {
		Name string `json:"name"`
		Role string `json:"role"`
	}

	for i := 0; i < b.N; i++ {
		A := &a{Name: fmt.Sprintf("name%v", i)}
		B := &s{}

		marshall.Marshall(A,B)
		b.Log("A: ", A)
		b.Log("B: ", B)
	} 
}