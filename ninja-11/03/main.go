package main

import "fmt"

type customErr struct {
	msg string
}

func (c customErr) Error() string {
	return c.msg
}

func main() {
	e := customErr{"This is a made up error"}
	foo(e)
}

func foo(e error) {
	fmt.Println("Foo ran")
	fmt.Println(e)
	fmt.Printf("%T\n", e)
}
