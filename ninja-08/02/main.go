package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First, Last string
	Age         int
	Sayings     []string
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	var people []person

	if err := json.Unmarshal([]byte(s), &people); err != nil {
		fmt.Printf("There was an error:\n%q\n", err)
	}

	for _, p := range people {
		fmt.Printf("%v %v, %d y/o likes to say\n", p.First, p.Last, p.Age)
		for _, s := range p.Sayings {
			fmt.Printf("\t%q\n", s)
		}
		fmt.Println()
	}

}
