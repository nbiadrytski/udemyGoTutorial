package main

import "fmt"

func main() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentValue := range greeting {
		fmt.Println(i, currentValue)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}
