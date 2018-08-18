package main

import "fmt"

func main() {
	var smallNumber int
	var largeNumber int

	fmt.Println("Enter a small number")
	fmt.Scan(&smallNumber)

	fmt.Println("Enter a large number")
	fmt.Scan(&largeNumber)

	remainder := largeNumber % smallNumber
	fmt.Println(remainder)

}
