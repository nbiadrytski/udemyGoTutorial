package main

import "fmt"

func main() {
	a := 43
	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a) //in hex
	fmt.Printf("%d \n", &a)                  //in decimal
}
