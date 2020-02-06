package main

import "fmt"

func main() {
	mInt := 340
	fmt.Printf("%d\t%b\t%x\n", mInt, mInt, mInt)
	shifted := mInt << 1
	fmt.Printf("%d\t%b\t%x\n", shifted, shifted, shifted)
}
