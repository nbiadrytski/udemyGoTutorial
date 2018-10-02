package main

import (
	"fmt"
	"math"
)

type square struct { // user defined type square, and struct is its underlying type
	side float64
}

// another shape
type circle struct {
	radius float64
}

// interface defines functionality
type shape interface { // user defined type shape, and interface is its underlying type
	area() float64
}

// concrete type square implements shape interface because it has this method signature "area() float64"
func (z square) area() float64 { // attaching method to square type
	return z.side * z.side
}

// concrete type circle implements shape interface because it has this method signature "area() float64"
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) { // you can pass anything that imlements type shape
	fmt.Println(z)
	fmt.Println(z.area())
}
func main() {
	s := square{10}
	c := circle{5}
	info(s)
	info(c)

}
