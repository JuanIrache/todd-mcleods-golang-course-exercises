package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side int
}

func (s square) area() float64 {
	return float64(s.side) * float64(s.side)
}

type circle struct {
	radius int
}

func (c circle) area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

func info(s shape) {
	fmt.Printf("The area of this shape is %f\n", s.area())
}

func main() {
	s := square{10}
	c := circle{10}
	info(s)
	info(c)
}
