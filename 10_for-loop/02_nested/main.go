package main

import "fmt"

func main()  {
	fmt.Println("i  -  j")
	for i :=0; i < 5; i++ {
		for j := 0; j < 5; j++{
			fmt.Println(i, " - ", j)
		}
	}
}
