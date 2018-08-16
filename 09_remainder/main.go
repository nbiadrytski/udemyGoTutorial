package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			fmt.Println(i, "is odd")
		} else {
			fmt.Println(i, "is even")
		}
	}

	var userNumber int

	fmt.Println("Enter number:")
	fmt.Scan(&userNumber)

	switch userNumber % 2 {
	case 1:
		fmt.Println(userNumber, "is odd")
	case 0:
		fmt.Println(userNumber, "is even")
	default:
		fmt.Println(userNumber, "???")
	}
}
