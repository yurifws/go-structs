package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	jim := person{
		firstName: "Jim",
		lastName:  "Carrey",
		contactInfo: contactInfo{
			email:   "jimcarrey@gmail.com",
			zipCode: 45126789,
		},
	}

	fmt.Printf("%+v", jim)

}
