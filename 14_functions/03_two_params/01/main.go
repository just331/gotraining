package main

import "fmt"

func main()  {
	greet("Justin", "Cennah")
}

func greet(fname string, sname string) {
	fmt.Println(fname,sname)
}