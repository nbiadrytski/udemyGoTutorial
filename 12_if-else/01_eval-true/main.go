package main

import "fmt"

func main() {

	if true {
		fmt.Println("This ran")
	}

	if false {
		fmt.Println("This did not run")
	}

	if !true {
		fmt.Println("This did not run !true")
	}

	if !false {
		fmt.Println("This ran !false")
	}
}