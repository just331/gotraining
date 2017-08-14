package main

import (
	"fmt"
	"sort"
)

func main()  {
	n := []int{9,4,5,7,3,8,32,66,49}

	fmt.Println(n)
	sort.Sort(sort.IntSlice(n))
	fmt.Println(n)
}