package main

import "fmt"

func one() {
	fmt.Println("one")
}

func two() {
	fmt.Println("two")
}

func three() {
	fmt.Println("three")
}

func main() {
	defer one() // func one() will be executed the last one
	two()
	three()
}
