package main

import "fmt"

func main() {
	xxs := make([][]string, 3)
	xxs[0] = []string{"First", "last", "Quote"}
	xxs[1] = []string{"James", "Bond", "Shaken, not stirred"}
	xxs[2] = []string{"Miss", "Penny", "Hello, James"}
	for _, xs := range xxs {
		for _, v := range xs {
			fmt.Printf("%q\t\t", v)
		}
		fmt.Print("\n")
	}
}
