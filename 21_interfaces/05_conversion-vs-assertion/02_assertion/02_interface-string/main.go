package main

import "fmt"

func main() {
	var name interface{} = "Nick" // value 7 would print "value is not a string"
	str, ok := name.(string)      // asserting that name is string
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
