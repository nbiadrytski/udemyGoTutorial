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

	buckets := make([]int, 300)

	for scanner.Scan() {
		n := hashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123])

}

func hashBucket(word string) int {
	return int(word[0])
}
