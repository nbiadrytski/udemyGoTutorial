package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	side float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.side * s.side
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, v := range shapes {
		area += v.area()
	}
	return area
}

func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)
	fmt.Println("Toatal area:", totalArea(c, s))
}