package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good Morning!",
		1: "Bonjour!",
		2: "Howdy!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 3)
	fmt.Println(myGreeting)
}
