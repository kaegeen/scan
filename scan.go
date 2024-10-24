package main

import "fmt"

func main() {
	var num1, num2 int

	// Input from the user
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Calculating the sum
	sum := num1 + num2

	// Displaying the result
	fmt.Printf("The sum of %d and %d is %d\n", num1, num2, sum)
}
