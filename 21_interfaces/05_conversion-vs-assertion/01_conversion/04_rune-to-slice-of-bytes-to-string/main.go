package main

import "fmt"

func main() {
	fmt.Println([]byte{'h', 'e', 'l', 'l', 'o'})         // [104 101 108 108 111]
	fmt.Println(string([]byte{'h', 'e', 'l', 'l', 'o'})) // hello
	// conversion: []bytes to string
}
