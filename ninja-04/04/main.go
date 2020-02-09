package main

import "fmt"

func main() {
	x := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	x = append(x, 0)
	x = append(x, -1, -2, -3)
	y := []int{-4, -5, -6, -7, -8, -9}
	x = append(x, y...)
	fmt.Println(x)
}
