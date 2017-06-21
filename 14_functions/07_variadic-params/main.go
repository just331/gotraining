package main

import "fmt"

func main(){
	n := average(45, 47, 36, 48, 73, 311)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf{
		total += v
	}
	return total / float64(len(sf))
}