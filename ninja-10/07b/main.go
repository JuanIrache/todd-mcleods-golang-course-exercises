package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan float64)
	full := make(chan int)
	const routines = 10
	const valuesPerRoutine = 10
	for i := 0; i < routines; i++ {
		go func() {
			for i := 0; i < valuesPerRoutine; i++ {
				c <- rand.Float64()
			}
			full <- 1
		}()
	}

	for i := 0; i < routines*valuesPerRoutine; i++ {
		fmt.Printf("Random float is %f ", <-c)
	}
	fmt.Printf("\n\nQuitting")
}
