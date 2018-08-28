package main

import "fmt"

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	fmt.Println(half(5))
	fmt.Println(half(8))
}

func half(x int) (divisionResult int, isEven bool) {
	divisionResult = x / 2
	isEven = x%2 == 0
	return divisionResult, isEven
}
