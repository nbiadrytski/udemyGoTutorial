package main

import (
	"fmt"
)

func main() {
	// slice has an underlying array,
	// but it's lenth is 0
	// that's why student == nil ---> false
	student := []string{}
	students := [][]string{}

	// length of underlying array is not set, basically it is 0
	// that's why student[0] = "Nick" will throw 'index out of range'
	// but student = append(student, "Nick") will work

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	/*
		[]
		[]
		false
	*/
}
