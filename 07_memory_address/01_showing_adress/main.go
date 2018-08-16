package main

import "fmt"

func main() {
	a := 43
	b := &a
	fmt.Println("this is the value of a:", a, ".", "this is a's address in hex:", b)
	fmt.Println(a, *b)
	fmt.Println(&a, &b)

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a) //in hex
	fmt.Printf("%d \n", &a)                  //in decimal
}
