package main

import "fmt"

func main() {
	arr := [5]int{5, 4, 3, 2, 1}
	fmt.Printf("Looping through range of array of type %T\n", arr)
	for _, v := range arr {
		fmt.Printf(" %d", v)
	}
}
