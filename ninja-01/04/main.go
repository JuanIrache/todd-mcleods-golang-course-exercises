package main

import "fmt"

type mInt int

var x mInt

func main() {
	fmt.Println("x:", x)
	fmt.Printf("x is of type %T\n", x)
	x = 42
	fmt.Println("new x:", x)
}
