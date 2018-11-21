package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		// Println formats using the default formats for its operands and writes to standard output.
		// err happened (fmt.Println) open no-file.txt The system cannot find the file specified
		fmt.Println("err happened (fmt.Println)", err)

		// same as above, but also prints date and time
		// 2018/11/21 09:11:44 err happened open no-file.txt: The system cannot find the file specified.
		log.Println("err happened (log.Println)", err)

		// 2018/11/21 09:16:55 err happened (log.Fatalln) open no-file.txt: The system cannot find the file specified.
		// exit status 1
		log.Fatalln("err happened (log.Fatalln)", err)

		// panic: open no-file.txt: The system cannot find the file specified.
		// goroutine 1 [running]:
		// main.main()
		// C:/Software/go/src/github.com/nbiadrytski/udemyGoTutorial/23_error-handling/02_err-not-nil/01_err-fmt-println/main.go:22 +0x6a
		// exit status 2
		panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output.

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.

/*
Package log implements a simple logging package ...
writes to standard error and prints the date and time of each logged message ...
the Fatal functions call os.Exit(1) after writing the log message ...
the Panic functions call panic after writing the log message.
*/

// log.Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
