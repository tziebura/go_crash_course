package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string)

	// Assign kv
	// emails["Bob"] = "bob@example.com"
	// emails["Sharon"] = "sharon@example.com"

	// Define map and assign kv
	emails := map[string]string{"Bob": "bob@example.com", "Sharon": "sharon@example.com"}

	fmt.Printf("Size of map: %d\n", len(emails))
	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Printf("Size of map: %d\n", len(emails))
	fmt.Println(emails)
}
