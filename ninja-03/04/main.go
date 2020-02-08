package main

import "fmt"

func main() {
	year := 1985
	for {
		if year == 2020 {
			fmt.Printf("%d", year)
			break
		}
		fmt.Printf("%d, ", year)
		year++
	}
}
