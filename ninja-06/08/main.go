package main

import (
	"fmt"
	"math/rand"
)

func main() {
	f := outer()
	fmt.Printf("Generated random float is %f\n", f())
}

func outer() func() float64 {
	r := rand.Float64()
	return func() float64 {
		return r
	}
}
