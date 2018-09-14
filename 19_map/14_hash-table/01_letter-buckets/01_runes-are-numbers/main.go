package main

import "fmt"

func main() {
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	// string to rune conversion
	letter = rune("A"[0])
	fmt.Println(letter)
}
