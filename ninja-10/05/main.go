package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 11
	}()

	v, ok := <-c
	fmt.Println("Channel is open. Value:", v, "Ok:", ok)

	close(c)

	v, ok = <-c
	fmt.Println("Channel was closed. Value:", v, "Ok:", ok)
}
