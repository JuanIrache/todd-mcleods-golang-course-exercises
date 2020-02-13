package main

import "fmt"

func main() {
	n := 40
	fmt.Printf("Fibonacci %d is %d\n", n, fibonacci(n))
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
