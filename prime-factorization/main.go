package main

import "fmt"

func main() {
	// Find all prime factors of a given number

	// Get user input
	fmt.Println("Enter a number: ")
	var input int
	fmt.Scanf("%d", &input)

	// Create a slice of integers to store the prime numbers
	var primeFactors []int

	// Get the number of 2s that divide the input
	for input%2 == 0 {
		primeFactors = append(primeFactors, 2)
		// Divide the input by 2
		input /= 2
	}

	// After finding the number of 2s, loop through the numbers from 3 because remaining input are odd
	for i := 3; i*i <= input; i = i + 2 {
		// If i divides input , print i and divide input by i
		for input%i == 0 {
			primeFactors = append(primeFactors, i)
			input /= i
		}
	}

	// Handle the case when input is a prime number greater than 2
	if input > 2 {
		primeFactors = append(primeFactors, input)
	}

	// Print the prime factors
	fmt.Println(primeFactors)

}
