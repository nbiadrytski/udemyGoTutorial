package main

import (
	"fmt"
)

func main() {
	var records [][]string

	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)

	// student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// store the record
	records = append(records, student2)

	// student 3
	student3 := make([]string, 4)
	student3[0] = "Nick"
	student3[1] = "Marriot"
	student3[2] = "32.00"
	student3[3] = "16.00"
	// store the record
	records = append(records, student3)

	fmt.Println(records)
}
