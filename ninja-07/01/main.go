package main

import "fmt"

func main() {
	s := "This is stored somewhere in memory"
	fmt.Printf("The address of s is %#x", &s)
}
