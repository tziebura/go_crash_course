package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "John", lastName: "Doe", city: "London", gender: "m", age: 30}
	person2 := Person{firstName: "Jane", lastName: "Smith", city: "London", gender: "f", age: 22}

	// Alternatvie
	// person2 := Person{"John", "Doe", "London", "m", 30}
	// fmt.Println(person1)
	// fmt.Println(person2)

	// person1.age++
	// fmt.Println(person1)

	person1.hasBirthday()
	person2.getMarried(person1.lastName)
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
