package main

import "fmt"

func main() {
	x := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	x = append(x[:3], x[6:]...)
	fmt.Println(x)
}
