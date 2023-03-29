package main

import "fmt"

func main() {
	// Using var
	// var name string = "John"
	var age int = 37
	const isCool = true

	// Using shorthand
	// name := "John"
	// size := 1.3

	name, size := "John", 1.3

	fmt.Println(name, age, isCool)
	fmt.Printf("%T", size)
}
