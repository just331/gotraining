package main

import "fmt"

func main() {
	var num1 int
	var num2 int
	fmt.Println("Enter a small number: ")
	fmt.Scan(&num1)
	fmt.Println("Enter a large number: ")
	fmt.Scan(&num2)
	//fmt.Println("The remainder of num2 and num1 is:")
	//fmt.Println(num2 % num1) my idea

	fmt.Println(num2, "%", num1, " = ", num2&num1)
}
