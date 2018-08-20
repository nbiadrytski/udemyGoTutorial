package main

import "fmt"

func main() {
	greeting := func() { // anonymous func; assign func to variable -> called func expression
		fmt.Println("Hello world!")
	}

	greeting()
	fmt.Printf("%T\n", greeting)
}
