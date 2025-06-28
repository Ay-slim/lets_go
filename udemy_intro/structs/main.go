package main

import "fmt"

type person struct {
	firstname string
	lastname string
	data contactInfo
}

type contactInfo struct {
	email string
	zipcode int
}

func main () {
	ayo := person {
		firstname: "Ayooluwa",
		lastname: "Adedipe",
		data: contactInfo {
			email: "perindo07@gmail.com",
			zipcode: 12345,
		},
	}
	ayo.print()
	ayo.updateFirstName("Isaiah")
	ayo.print()
	var ese person
	ese.print()
	ese.firstname = "Eseosa"
	ese.lastname = "Azenabor"
	ese.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateFirstName (newFirstName string) {
	p.firstname = newFirstName
}