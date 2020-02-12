package main

import "fmt"

func main() {
	var xi []int
	for i := 0; i <= 1000; i++ {
		xi = append(xi, i)
	}
	fmt.Printf("Total of xi by variadic is %d\n", foo(xi...))
	fmt.Printf("Total of xi by slice is %d\n", bar(xi))
}

func foo(xi ...int) int {
	var total int
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	var total int
	for _, v := range xi {
		total += v
	}
	return total
}
