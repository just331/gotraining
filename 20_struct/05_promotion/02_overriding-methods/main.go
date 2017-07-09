package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regualr person")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss MoneyPenny, so good to see you!")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  33,
	}

	p2 := doubleZero{

		person: person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()

}
