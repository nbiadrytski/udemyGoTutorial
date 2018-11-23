package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		errCustom := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, errCustom
	}
	// implementation
	return 42, nil
}
