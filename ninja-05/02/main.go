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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		printPerson(v)
		print("\n\n")
	}
}

func printPerson(p person) {
	fmt.Printf("First name:\t %q\n", p.first)
	fmt.Printf("Last name:\t %q\n", p.last)
	fmt.Print("Flavours:\t")
	for _, v := range p.flavours {
		fmt.Printf(" %q", v)
	}
}
