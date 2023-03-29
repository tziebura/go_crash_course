package main

import "fmt"

func adder() func(int) int {
	sum := 0
	fmt.Println("Sum is equal to:", sum) // Executed once
	return func(x int) int {
		sum += x
		fmt.Println("Sum in nested func is equal to:", sum) // Executed n times
		return sum
	}
}

func main() {
	adder := adder()
	fmt.Printf("%T\n", adder)
	for i := 0; i < 10; i++ {
		fmt.Println(adder(i))
	}
}
