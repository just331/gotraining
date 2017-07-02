package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good Morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Justin"] = "Howdy!"
	fmt.Println(myGreeting)
	myGreeting["Justin"] = "Gidday"
	fmt.Println(myGreeting)

}
