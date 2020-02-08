package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("x: %v\ty: %v\tz: %v", x, y, z)
	fmt.Print(s)
}
