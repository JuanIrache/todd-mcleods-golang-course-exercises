package main

import (
	"fmt"
	"math/rand"
)

func main() {
	c := gen()
	for v := range c {
		fmt.Printf("Random number: %d ", v)
	}
	fmt.Println("\nExiting")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- rand.Intn(100)
		}
		close(c)
	}()
	return c
}
