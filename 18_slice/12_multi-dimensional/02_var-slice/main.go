package main

import (
	"fmt"
)

func main() {
	// var sets the slice to the 0 value
	// the 0 value of slice is nil
	var student []string
	var students [][]string

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	/*
		[]
		[]
		true
	*/
}
