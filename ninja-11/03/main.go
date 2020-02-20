package main

import (
	"errors"
	"fmt"
)

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
	fmt.Println("e can use assertion to use as customErr and retrieve msg: ", e.(customErr).msg)
	y := errors.New("New error")
	if msg, ok := y.(customErr); ok {
		fmt.Println("new error y can use assertion to use as customErr and retrieve msg: ", msg)
	} else {
		fmt.Println("new error y can not use assertion to use as customErr and retrieve msg: ")

	}
}
