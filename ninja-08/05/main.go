package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (ba byAge) Less(i, j int) bool {
	return ba[i].Age < ba[j].Age
}

func (ba byAge) Len() int {
	return len(ba)
}

func (ba byAge) Swap(i, j int) {
	ba[j], ba[i] = ba[i], ba[j]
}

type byLast []user

func (bl byLast) Less(i, j int) bool {
	return bl[i].Last < bl[j].Last
}

func (bl byLast) Len() int {
	return len(bl)
}

func (bl byLast) Swap(i, j int) {
	bl[j], bl[i] = bl[i], bl[j]
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	uu := []user{u1, u2, u3}

	for _, u := range uu {
		sort.Strings(u.Sayings)
	}

	sort.Sort(byAge(uu))
	fmt.Println("By Age")
	printUsers(uu)

	sort.Sort(byLast(uu))
	fmt.Println("By Last")
	printUsers(uu)

}

func printUsers(uu []user) {
	for _, u := range uu {
		fmt.Printf("%v %v, %d y/o likes to say\n", u.First, u.Last, u.Age)
		for _, s := range u.Sayings {
			fmt.Printf("\t%q\n", s)
		}
		fmt.Println()
	}
}
