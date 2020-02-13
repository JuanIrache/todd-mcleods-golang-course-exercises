package main

import "fmt"

type person struct {
	name string
	age  int
}

func changeMe(p *person) {
	p.age++
}

func main() {
	p := person{"John", 10}
	fmt.Printf("%v is %d years old\n", p.name, p.age)
	changeMe(&p)
	fmt.Printf("%v is %d years old\n", p.name, p.age)
}
