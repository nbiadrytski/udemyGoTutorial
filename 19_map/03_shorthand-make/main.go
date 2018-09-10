package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	myGreeting["Nick"] = "Good morning"
	myGreeting["John"] = "Hello"

	fmt.Println(myGreeting)
}
