package main

import "fmt"

func main() {
	f := func() {
		fmt.Print("This is being called from within a variable")
	}
	f()
}
