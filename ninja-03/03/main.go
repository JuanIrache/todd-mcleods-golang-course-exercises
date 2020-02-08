package main

import "fmt"

func main() {
	year := 1985
	for year < 2020 {
		fmt.Printf("%d, ", year)
		year++
	}
	fmt.Printf("%d", 2020)
}
