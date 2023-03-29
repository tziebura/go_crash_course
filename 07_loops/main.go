package main

import "fmt"

func main() {
	i := 1

	// Long method
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short method
	for i := 1; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Printf("%d: FizzBuzz \n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d: Buzz \n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d: Fizz \n", i)
		} else {
			fmt.Println(i)
		}
	}
}
