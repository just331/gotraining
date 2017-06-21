package main

import "fmt"

func main()  {
	fmt.Println(greet("Justin ", "Cennah"))
}

func greet(fname,sname string) (s string)  {
	s = fmt.Sprint(fname,sname)
	return
}


/*
IMPORTANT
Avoid using named returns

Occasionally named returns are useful.

 */