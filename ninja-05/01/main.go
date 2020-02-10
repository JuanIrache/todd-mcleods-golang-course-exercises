package main

import "fmt"

type person struct {
	first, last string
	flavours    []string
}

func main() {
	p1 := person{
		first:    "John",
		last:     "Doe",
		flavours: []string{"strawberry", "peppermint", "pistachio"},
	}
	p2 := person{
		first:    "Michael",
		last:     "Collins",
		flavours: []string{"beer", "coffee", "crisps"},
	}
	printPerson(p1)
	fmt.Print("\n\n")
	printPerson(p2)
}

func printPerson(p person) {
	fmt.Printf("First name: %q\n", p.first)
	fmt.Printf("Last name: %q\n", p.last)
	fmt.Print("Flavours:")
	for _, v := range p.flavours {
		fmt.Printf(" %q", v)
	}
}
