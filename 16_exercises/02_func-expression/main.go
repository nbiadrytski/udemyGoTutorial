package main

import "fmt"

func main() {

	half := func(x int) (divisionResult int, isEven bool) {
		divisionResult = x / 2
		isEven = x%2 == 0
		return divisionResult, isEven
	}
	fmt.Println(half(5))
}
