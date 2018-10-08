package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	// sorting user-defined type
	studyGroup := people{"Zeno", "Joe", "Al", "Jenny"}
	fmt.Println("before sorting (user-defined type)", studyGroup)

	sort.Sort(studyGroup)
	fmt.Println("after sorting (user-defined type)", studyGroup)

	sort.Sort(sort.Reverse(studyGroup))
	fmt.Println("after reverse (user-defined type)", studyGroup)

	// sorting slice of strings
	s := []string{"Zeno", "Joe", "Al", "Jenny"}
	fmt.Println("before sorting (slice of strings)", s)
	// sort.Strings(s)
	// another way; turns s into StringSlice type which has Sort(), Swap(), etc. methods
	sort.StringSlice(s).Sort()
	fmt.Println("after sorting (slice of strings)", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("after reverse (slice of strings)", s)

	// sorting slice of ints
	n := []int{5, 2, 6, 3, 1, 4}
	fmt.Println("before sorting (slice of ints)", n)
	//sort.Ints(n)
	// another way; turns s into IntSlice type which has Sort(), Swap(), etc. methods
	sort.IntSlice(n).Sort()
	fmt.Println("after sorting (slice of ints)", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("after reverse (slice of ints)", n)
}
