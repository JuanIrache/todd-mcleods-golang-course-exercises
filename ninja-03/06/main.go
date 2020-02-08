package main

import "fmt"

func main() {
	this := true
	if this {
		that()
	}
}

func that() {
	fmt.Printf("That")
}
