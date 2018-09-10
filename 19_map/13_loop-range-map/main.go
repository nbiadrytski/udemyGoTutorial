package main

import "fmt"

func main() {
	myGretting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	for key, value := range myGretting {
		fmt.Println(key, " - ", value)
	}
}
