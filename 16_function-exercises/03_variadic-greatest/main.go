package main

import "fmt"



func main()  {

	max := greatest(1,5,8,10,45,700,44,44,110,5,500)
	fmt.Println(max)

}

func greatest(g ...int) int {
 var largest int
	for _, v := range g {
		if v > largest{
			largest = v
		}
	}
	return largest
}

