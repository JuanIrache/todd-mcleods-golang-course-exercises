package main

import "fmt"

var favSport string = "fencing"

func main() {
	switch favSport {
	case "soccer":
		fmt.Println("Booh!")
	case "fencing":
		fmt.Println("Rocks!")
	}
}
