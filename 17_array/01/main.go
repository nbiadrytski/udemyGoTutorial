package main

import "fmt"

func main() {
	var x [58]string // initialises things to 0 for int and '' for string

	fmt.Println("array before initialisation:", x)
	fmt.Println("array length before initialisation:", len(x))
	fmt.Println("array's index 42 before initialisation:", x[42])

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	fmt.Println("array after initialisation:", x)
	fmt.Println("array length after initialisation:", len(x))
	fmt.Println("array's index 42 after initialisation:", x[42])

	fmt.Println(string(65))
}
