package main

import "fmt"

func main() {
	for i := 65; i < 91; i++ {
		fmt.Printf("%d\t", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%#U\t", i)
		}
		fmt.Print("\n")
	}
}
