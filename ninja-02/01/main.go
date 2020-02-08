package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := 1234
	//Addition: Can read the first number of the command as an int
	if len(os.Args) > 1 {
		n, _ = strconv.Atoi(os.Args[1])
	}
	fmt.Printf("%v\t%b\t%x", n, n, n)
}
