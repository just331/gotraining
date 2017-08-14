package main

import (
	"sort"
	"fmt"
)

func main() {
	s := []string{"Xeno", "Jenny", "Hayli", "Al"}
	sort.Strings(s)
	fmt.Println(s)
}