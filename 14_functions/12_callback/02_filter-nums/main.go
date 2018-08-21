package main

import "fmt"

func filter(numbers []int, callbackFunc func(int) bool) []int { // this is type of callbackFunc func: ----> func(int) bool
	var xs []int
	for _, n := range numbers {
		if callbackFunc(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T", callbackFunc)
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool { // the whole func passed as an argument
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]
}
