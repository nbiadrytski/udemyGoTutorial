package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://data.iana.org/TLD/tlds-alpha-by-domain.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(resp.Body)
	str := string(bs)

	fmt.Println(str)
}
