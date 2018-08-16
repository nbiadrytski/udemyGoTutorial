package main

import "fmt"

func zero(x *int) { //func param is memory address of x
	fmt.Println("x address in func zero", x)
	*x = 0 // setting the value of memory address to 0
}

func main() {
	x := 5
	fmt.Println("x address in func main", &x)
	zero(&x)                        // passing 5 to memory address
	fmt.Println("final x value", x) // x is 0
}
