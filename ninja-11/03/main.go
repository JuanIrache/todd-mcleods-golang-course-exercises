package main

import "fmt"

type customErr struct {
	msg string
}

func (c customErr) Error() string {
	return fmt.Sprint(c.msg)
}

func main() {
	e := customErr{"This is a made up error"}
	foo(e)
}

func foo(e error) {
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(e.(error))
}
