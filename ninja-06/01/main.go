package main

import "fmt"

func main() {
	fmt.Printf("foo returns %d\n", foo())
	d, q := bar()
	fmt.Printf("bar returns %d and %q\n", d, q)
}

func foo() int {
	return 10
}

func bar() (int, string) {
	return 10, "ten"
}
