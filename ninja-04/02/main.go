package main

import "fmt"

func main() {
	x := []int{10, 9, 8, 7, 6, 6, 5, 4, 3, 2, 1}
	fmt.Printf("Looping through range or slice of type %T\n", x)
	for _, v := range x {
		fmt.Printf(" %d", v)
	}
}
