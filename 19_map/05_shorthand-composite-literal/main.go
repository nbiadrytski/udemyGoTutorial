package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good morning",
		"Jenny": "Hello",
	}
	fmt.Println(myGreeting["Jenny"])
	// Hello

	myGreeting["Teddy"] = "Hi"
	fmt.Println(myGreeting)
	// map[Tim:Good morning Jenny:Hello Teddy:Hi]

	myGreeting["Teddy"] = "Привет"
	fmt.Println(myGreeting)
	// map[Tim:Good morning Jenny:Hello Teddy:Привет]

	fmt.Println(len(myGreeting))

	delete(myGreeting, "Teddy")
	fmt.Println(myGreeting)
	// map[Tim:Good morning Jenny:Hello]
}
