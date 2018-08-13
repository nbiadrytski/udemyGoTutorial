package main

import "fmt"

// the scope of x is this whole package
var x = 42

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}
