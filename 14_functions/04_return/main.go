package main

import "fmt"

func main()  {
	fmt.Println(greet("Justin ", "Doe"))
}

func greet(fname,sname string) string  {
	return fmt.Sprint(fname,sname)
}