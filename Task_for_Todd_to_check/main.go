package main

import (
	"fmt"
	"time"
)

func main() {
	var sundays int
	var month time.Month
	i := int(month)
	for year := 1901; year <= 2000; year++ {
		for i = 1; i <= 12; i++ {
			t := time.Date(year, time.Month(i), 1, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == 6 {
				sundays++
			}
		}
	}
	fmt.Println(sundays, " Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)")
}

/*
https://projecteuler.net/problem=19
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
*/
