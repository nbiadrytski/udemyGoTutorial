package main

import "fmt"

const p = "death & taxes" //untyped constant

func main() {
	const q int = 42 //typed constant

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
}

// a CONSTANT is a simple unchanging value
