package main

import "fmt"

type mString string

var x mString

var y string

func main() {
	fmt.Println("x:", x)
	fmt.Printf("x is of type %T\n", x)
	x = "not zero"
	fmt.Println("New x:", x)
	y = string(x)
	fmt.Println("y:", y)
	fmt.Printf("y is of type %T\n", y)
}
