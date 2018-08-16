package main

import "fmt"

func zero(x int) {
	fmt.Println("x address in func zero", &x) // x address in func zero 0xc420012070
	x = 0
	fmt.Println("x from zero func", x) // x = 0
}

func main() {
	x := 5
	fmt.Println("x address in func main", &x) // x address in func main 0xc420012058
	zero(x)
	fmt.Println("x from main func", x) // x is still 5; we passed the value of x (which is 5), not the memory address to func zero
	// that's why x is not changed by func zero
}
