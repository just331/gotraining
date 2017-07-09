package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{"Justin", "Rodriguez", 20}
	p2 := person{"Miss", "Moneypenny", 21}

	fmt.Printf(p1.first, p1.last, p1.age)
	fmt.Printf(p2.first, p2.last, p2.age)
}
