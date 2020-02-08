package main

import "fmt"

func main() {
	switch {
	case "Life is a lie" == "true":
		fmt.Println("This is an error")
	case "I look like the string next to me" == "I look like the string next to me":
		fmt.Println("I do look like the string next to me")
	}
}
