package main

import (
	"fmt"

	"github.com/juanirache/todd-mcleods-golang-course-exercises/ninja-13/02/quote"
	"github.com/juanirache/todd-mcleods-golang-course-exercises/ninja-13/02/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
