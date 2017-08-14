package main

import (
	"fmt"
	"sort"
)

func main() {

	n:= []int{3,6,11,777,23,55,14234,666,333}

	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
