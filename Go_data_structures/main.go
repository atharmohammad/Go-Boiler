package main

import "fmt"

type ContactInfo struct {
	email    string
	password string
}

type Person struct {
	firstName string
	lastName  string
	contact   ContactInfo
}

func main() {
	person := Person{
		firstName: "Athar",
		lastName:  "Mohammad",
		contact: ContactInfo{
			email:    "athar@g.com",
			password: "1444",
		},
	}
	person.update("Anas")
	fmt.Println(person)
}

func (p *Person) update(newName string) {
	(*p).firstName = newName
}
