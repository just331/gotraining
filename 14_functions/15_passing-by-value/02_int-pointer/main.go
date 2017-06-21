package main

import "fmt"

func main()  {
	age := 44
	fmt.Println(&age) //0xc042006290

	changeMe(&age)

	fmt.Println(&age) //0xc042006290
	fmt.Println(age)  //44
}

func changeMe(z *int)  {
	fmt.Println(z) //0xc042006290
	fmt.Println(*z)
	*z = 24
	fmt.Println(z)
	fmt.Println(*z)
}