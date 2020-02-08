package main

import "fmt"

func main() {
	for i := 1; i < 10000; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Printf("%v", 10000)
}
