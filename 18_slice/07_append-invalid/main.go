package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "Buenos dias!"
	greeting[3] = "suprabadham!" // > than the length of the slice

	fmt.Println(greeting[2])

}
