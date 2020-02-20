package main

import "fmt"

func main() {
	fmt.Println(Factorial(10))
}

func Factorial(i int) int {
	if i == 1 {
		return i
	}
	return i * Factorial(i-1)
}
