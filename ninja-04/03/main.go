package main

import "fmt"

func main() {
	x := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
