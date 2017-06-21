package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please enter your nanme:")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
