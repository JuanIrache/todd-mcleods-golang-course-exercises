package main

import "fmt"

const thisYear = 2020

const (
	year1 = thisYear - iota
	year2 = thisYear - iota
	year3 = thisYear - iota
	year4 = thisYear - iota
)

func main() {
	fmt.Printf("%v\t%v\t%v\t%v", year1, year2, year3, year4)
}
