package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println("With a func literal", <-c)

	c2 := make(chan int, 1)

	c2 <- 42

	fmt.Println("With a buffered channel", <-c2)
}
