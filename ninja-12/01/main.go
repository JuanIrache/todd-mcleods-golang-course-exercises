package main

import (
	"fmt"

	"github.com/juanirache/todd-mcleods-golang-course-exercises/ninja-12/01/dog"
)

func main() {
	y := 3
	fmt.Printf("%d human years correspond to %d dog years\n", y, dog.YearsHumanToDog(y))
	y = 10
	fmt.Printf("%d dog years correspond to %d human years\n", y, dog.YearsDogToHuman(y))
}
