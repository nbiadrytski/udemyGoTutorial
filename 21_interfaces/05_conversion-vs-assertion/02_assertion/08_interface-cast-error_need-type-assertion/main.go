package main

import "fmt"

func main() {
	rem := 7.24
	fmt.Printf("%T\n", rem)
	fmt.Printf("this is casting %T\n", int(rem))

	var val interface{} = 7
	fmt.Printf("%T\n", val)
	fmt.Printf("this is assertion %T\n", val.(int))
}
