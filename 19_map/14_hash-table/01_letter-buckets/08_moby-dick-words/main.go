package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
