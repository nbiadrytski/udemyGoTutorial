package main

import "fmt"

func greatest(numbers []int) int {
	max := numbers[0]
	for _, value := range numbers {
		if value > max {
			max = value
		}
	}
	return max
}

func greatest2(number ...int) int {
	var max int
	for _, value := range number {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	var numbers1 = []int{2, 7, 9, 4, 89, 66, 12, 799, 999, 776543, 234}
	fmt.Println(greatest(numbers1))

	fmt.Println(greatest2(2, 5, 8, 9, 1, 44, 6))
}
