package main

import "fmt"

func main()  {
	fmt.Println(greet("Justin ","Rodriguez " ))
}

func greet(fname, sname string) (string,string) {
	return fmt.Sprint(fname,sname), fmt.Sprint(sname,fname)
}