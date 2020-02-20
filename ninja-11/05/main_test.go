package main

import "testing"

func TestFactorial(t *testing.T) {
	got := Factorial(10)
	if got != 3628800 {
		t.Errorf("Factorial(10) expected 3628800, got %d", got)
	}
}
