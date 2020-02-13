package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("The double of random is %d\n", runWithRandom(byTwo))
}

func byTwo(n int) int {
	return n * 2
}

func runWithRandom(f func(int) int) int {
	r := rand.Intn(100)
	return f(r)
}
