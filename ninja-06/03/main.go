package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Code on top")
	}()
	fmt.Println("Code on bottom")
}
