package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	if y := Years(2); y != 14 {
		t.Errorf("Years(2) expected 14, got %d", y)
	}
}

func TestYearsTwo(t *testing.T) {
	if y := Years(2); y != 14 {
		t.Errorf("Years(2) expected 14, got %d", y)
	}
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(2)
	}
}

func ExampleYears() {
	twoDogYears := Years(2)
	fmt.Println(twoDogYears)
	// Output: 14
}

func ExampleYearsTwo() {
	twoDogYears := YearsTwo(2)
	fmt.Println(twoDogYears)
	// Output: 14
}
