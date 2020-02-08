package main

import "fmt"

func main() {
	a := 1
	b := 2
	c := 2
	aEqualsB := a == b
	aSmallerOrEqualC := a <= c
	bLargerOrEqualC := b >= c
	bDifferentC := b != c
	cSmallerA := c < a
	cLargerA := c > a

	fmt.Println("does a equal b?", aEqualsB)
	fmt.Println("is a smaller or equal to c?", aSmallerOrEqualC)
	fmt.Println("is b larger or equal to c?", bLargerOrEqualC)
	fmt.Println("is b different than c?", bDifferentC)
	fmt.Println("is c smaller than a?", cSmallerA)
	fmt.Println("is c larger than a?", cLargerA)
}
