package main

import "fmt"

func main() {
	var myGreeting map[string]string
	fmt.Println(myGreeting)

	//myGreeting["Nick"] = "Good morning"
	// panic: assignment to entry in nil map

	fmt.Println(myGreeting == nil)
}
