package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("%d divided by 4 is %d and leaves a remainder of %d\n", i, i/4, i%4)
	}
}
