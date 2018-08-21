package main

import "fmt"

func visit(numbers []int, callbackFunk func(int)) { // callbackFunk is function name; func(int) --> it's a type
	for _, n := range numbers { // func visit() takes 2 params: a slice of int
		callbackFunk(n) // and takes a function which takes int
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

// callback: passing a func as an argument
