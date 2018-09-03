package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 3)

	fmt.Println("----------------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("----------------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), mySlice[i])
	}
}

/* When len = 3 and cap = 3
this means that underlying array is full;
that's why when we add new item to the slice, then
initial array with capacity 3 is deleted and
new array with doubled capacity 6 is created
*/
