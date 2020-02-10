package main

import "fmt"

func main() {
	complex := struct {
		integer  int
		floating float64
		text     string
		slice    []int
		kv       map[string]float64
	}{
		integer:  1,
		floating: 10.15,
		text:     "this is text",
		slice:    []int{9, 1, 1},
		kv: map[string]float64{
			"first":  1.0,
			"second": 2.0,
			"third":  3.0,
		},
	}
	fmt.Printf("Integer:\t%d\n", complex.integer)
	fmt.Printf("Floating:\t%f\n", complex.floating)
	fmt.Printf("Text:\t\t%q\n", complex.text)
	for i, v := range complex.slice {
		fmt.Printf("Slice value %d:\t%d\n", i, v)
	}
	for k, v := range complex.kv {
		fmt.Printf("Kv %q:\t%f\n", k, v)
	}
}
