package main

import "fmt"

// A B C D : 0, 1, 2, 3
const (
	A = iota
	B
	C
	D
)

const (
	e = iota
	f
	g
	h
)

func main() {
	fmt.Println("A", A)
	fmt.Println("B", B)
	fmt.Println("C", C)
	fmt.Println("D", D)

	fmt.Println("e", e)
	fmt.Println("f", f)
	fmt.Println("g", g)
	fmt.Println("h", h)

}
