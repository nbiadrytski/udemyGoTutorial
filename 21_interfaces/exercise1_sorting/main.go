package main

import (
	"fmt"
	"sort"
)

type people []string
type digits []int

func (p people) Len() int {
	return len(p)
}

func (d digits) Len() int {
	return len(d)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (d digits) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (d digits) Less(i, j int) bool {
	return d[i] < d[j]
}

func main() {
	studyGroup := people{"Zeno", "Joe", "Al", "Jenny"}
	fmt.Println("before sorting", studyGroup)

	sort.Sort(studyGroup)
	fmt.Println("after sorting", studyGroup)

	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println("after reverse", studyGroup)

	n := []int{2, 4, 3, 6}
	fmt.Println("before sorting", n)

	sort.Sort(digits(n))
	fmt.Println("after sorting", n)

	sort.Sort(sort.Reverse(digits(n)))
	fmt.Println("after reverse", n)

}
