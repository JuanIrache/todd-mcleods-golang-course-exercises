package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {

	vitara := truck{
		vehicle: vehicle{
			doors: 5,
			color: "red",
		},
		fourWheel: true,
	}
	jetta := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(vitara)
	fmt.Printf("Number of doors: %d\n", vitara.doors)
	fmt.Println()
	fmt.Println(jetta)
	fmt.Printf("Number of doors: %d\n", jetta.doors)

}
