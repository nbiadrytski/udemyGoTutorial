package main

import "fmt"

const (
	pi       = 3.14
	language = "Go"
	height   = 1.83
)

func main() {
	fmt.Println(pi)
	fmt.Println(language)
	fmt.Printf("%v - %T - %f", height, height, height)
}
