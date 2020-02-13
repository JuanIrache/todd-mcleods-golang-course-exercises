package main

import (
	"fmt"
	"math/rand"
)

func main() {
	mistery := randomEnclosedDoubler()
	fmt.Println(mistery())
	fmt.Println(mistery())
	fmt.Println(mistery())
	fmt.Println(mistery())
	fmt.Println(mistery())
}

func randomEnclosedDoubler() func() int {
	hiddenVar := rand.Intn(1000)
	return func() int {
		hiddenVar *= 2
		return (hiddenVar)
	}
}
