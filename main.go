package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

	sofia := person{"Sofia", "Lee"}

	fmt.Println(alex)
	fmt.Println(sofia)
}
