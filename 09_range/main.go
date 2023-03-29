package main

import "fmt"

func main() {
	ids := []int{33, 75, 1, 9}

	// Loop through ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Println("Sum equals:", sum)

	// Range with map
	emails := map[string]string{"Bob": "bob@example.com", "Sharon": "sharon@example.com"}

	for name, email := range emails {
		fmt.Printf("%s has email: %s\n", name, email)
	}

	for name := range emails {
		fmt.Println("Name: ", name)
	}
}
