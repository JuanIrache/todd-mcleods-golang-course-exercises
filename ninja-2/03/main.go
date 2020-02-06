package main

import "fmt"

const (
	typed   float64 = 10
	unTyped         = 10
)

func main() {
	fmt.Printf("Typed const is of type %T and value %v\n", typed, typed)
	fmt.Printf("Untyped const is of type %T and value %v\n", unTyped, unTyped)
}
