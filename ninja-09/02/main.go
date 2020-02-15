package main

import "fmt"

type person struct {
	name string
}

func (p *person) speak() {
	fmt.Printf("My name is %v\n", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p := person{"James Bond"}

	//Works
	saySomething(&p)
	//Works
	p.speak()
	//Does not work
	// saySomething(p)
}
