package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	xi := []int{0, 10, 4, 6}
	if got := CenteredAvg(xi); got != 5 {
		t.Errorf("CenteredAvg expected 5, got %f", got)
	}
}

type testMap struct {
	param  []int
	result float64
}

func TestCenteredAvg_table(t *testing.T) {
	table := []testMap{
		testMap{
			param:  []int{0, 10, 4, 6},
			result: 5,
		},
		testMap{
			param:  []int{1, 9, 5, 6},
			result: 5.5,
		},
		testMap{
			param:  []int{100, 101, 103, 200},
			result: 102,
		},
	}

	for _, v := range table {
		if got := CenteredAvg(v.param); got != v.result {
			t.Errorf("CenteredAvg table test expected %f got %f", v.result, got)
		}
	}

}

func BenchmarkCenteredAvg(b *testing.B) {
	xi := []int{0, 10, 4, 6}
	for i := 0; i < b.N; i++ {
		CenteredAvg(xi)
	}
}

func ExampleCenteredAvg() {
	xi := []int{0, 10, 4, 6}
	got := CenteredAvg(xi)
	fmt.Println(got)
	// Output: 5
}
