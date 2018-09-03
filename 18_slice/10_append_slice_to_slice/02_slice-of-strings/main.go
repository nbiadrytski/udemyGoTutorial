package main

import "fmt"

func main() {

	mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday"}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice, "before removal")

	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Println(mySlice, "after removal Wednesday")
}
