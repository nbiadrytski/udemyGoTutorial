package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://data.iana.org/TLD/tlds-alpha-by-domain.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(resp.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	i := 0
	for k := range words {
		fmt.Println(k)
		if i == 200 {
			break
		}
		i++

	}
}
