package word

import (
	"fmt"
	"testing"
)

func TestUseCount(t *testing.T) {
	s := "the type is all about the type"
	res := UseCount(s)
	if res["about"] != 1 || res["type"] != 2 {
		t.Error("UseCount did not return the correct values")
	}
}

func TestCount(t *testing.T) {
	s := "the type is all about the type"
	res := Count(s)
	if res != 7 {
		t.Errorf("Extected 7, got %d", res)
	}
}

func ExampleUseCount() {
	s := "the type is all about the type"
	res := UseCount(s)
	fmt.Println(res["type"])
	// Output: 2
}

func ExampleCount() {
	s := "the type is all about the type"
	res := Count(s)
	fmt.Println(res)
	// Output: 7
}

func BenchmarkUseCount(b *testing.B) {
	s := "the type is all about the type"
	for i := 0; i < b.N; i++ {
		UseCount(s)
	}
}

func BenchmarkCount(b *testing.B) {
	s := "the type is all about the type"
	for i := 0; i < b.N; i++ {
		Count(s)
	}
}
