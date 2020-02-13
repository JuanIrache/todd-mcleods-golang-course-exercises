package main

import "fmt"

func main() {
	func() {
		fmt.Print("This is an anonymous func")
	}()
}
