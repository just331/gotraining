package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Xeno", "John", "Jenny", "Hayli"}

	fmt.Println(s)
	// sort.StringSlice(S).Sort()
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
}


// https://golang.org/pkg/sort/#Sort