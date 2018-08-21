package main

import "fmt"

func changeMe(z int) {
	fmt.Println("z from changeMe()", z)
	z = 24
	fmt.Println("z from changeMe() after z = 24     ", z)
}

func main() {
	age := 44
	changeMe(age)
	fmt.Println("age from main() after changeMe() was called", age) // 44
}

// we are passing value, not the age
