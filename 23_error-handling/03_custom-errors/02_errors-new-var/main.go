package main

import (
	"errors"
	"fmt"
	"log"
)

var errCustom = errors.New("norgate math: square root of negative number")

func main() {
	fmt.Printf("%T\n", errCustom)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errCustom
	}
	// implementation
	return 42, nil
}
