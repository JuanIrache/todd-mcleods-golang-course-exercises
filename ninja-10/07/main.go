package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := make(chan float64)
	full := make(chan int)
	routines := 10
	for i := 0; i < routines; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- rand.Float64()
			}
			full <- 1
		}()
	}
	done := 0
	for {
		select {
		case v := <-c:
			fmt.Printf("Random float is %f ", v)
		case <-full:
			done++
			if done >= routines {
				fmt.Printf("\n\nQuitting")
				return
			}
		}
	}
}
