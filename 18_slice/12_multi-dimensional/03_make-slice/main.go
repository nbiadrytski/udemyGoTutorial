package main

import (
	"fmt"
)

func main() {
	// when you do make you set length and capacity
	// that's why student[0] = "Nick" will work

	// make is preferable way of creating a slice
	student := make([]string, 35)
	students := make([][]string, 35)

	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

	/*
			[                                  ]
		    [[] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] [] []]
		    false
	*/
}
