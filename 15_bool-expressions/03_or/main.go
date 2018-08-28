package main

import "fmt"

func main() {
	if true || false {
		fmt.Println("This ran")
	}

	if true && false {
		fmt.Println("This didn't run")
	}
}
