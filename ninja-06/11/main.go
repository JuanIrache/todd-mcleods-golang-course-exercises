package main

import "fmt"

func main() {
	n := 100
	fmt.Printf("Fibonacci %d is %d\n", n, fibonacci(n))
}

var memo = map[int]int{
	0: 0,
	1: 1,
}

func fibonacci(n int) int {
	if _, ok := memo[n]; !ok {
		memo[n] = fibonacci(n-1) + fibonacci(n-2)
	}
	return memo[n]
}
