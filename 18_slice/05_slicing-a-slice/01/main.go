package main

import "fmt"

func main() {

	var results []int
	fmt.Println(results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)

	fmt.Println(mySlice[2:4]) //slicing a slice; c, g

	fmt.Println(mySlice[2]) // index access; accessing by index; c

	fmt.Println("myString"[2]) // index access; accessing by index; index 2 is S which is 83 in ASCII

	/* string is a slice of bytes
	string is made up of runes
	rune is a unicode code point
	a unicode code point is 1 to 4 bytes
	*/

	/* these are the same
	make([]int, 50, 100)
	new([100]int)[0:50]
	*/
}
