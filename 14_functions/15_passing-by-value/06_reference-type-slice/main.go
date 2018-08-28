package main

import "fmt"

func main() {
	m := make([]string, 1, 25)                                       // create a slice of strings, lenght 1, capacity 25
	fmt.Println("slice is empty as changeMe() wasn't called yet", m) // []
	m[0] = "Nick"
	fmt.Println("m[0] is now ", m) // []
	changeMe(m)
	fmt.Println("m[0] is now ", m) // [Todd]

}

func changeMe(z []string) { // z is a slice param
	z[0] = "Todd"
	fmt.Println("this value is from changeMe()", z) // [Todd]
}

/*
Slice is a reference type
We pass the value of the reference (which is the address actually) and it changes the value in that address
that's why the inintial value is changed after the function call
*/
