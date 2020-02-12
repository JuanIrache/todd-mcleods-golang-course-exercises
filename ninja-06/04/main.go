package main

import "fmt"

type person struct {
	first, last string
	age         int
}

func (p person) speak() {
	fmt.Printf("I'm %v %v and I'm %v years old\n", p.first, p.last, p.age)
}

func main() {
	p := person{
		"Juan",
		"Irache",
		34,
	}

	p.speak()
}
