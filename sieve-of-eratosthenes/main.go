// Print the prime numbers up to user input using the Sieve of Eratosthenes algorithm
package main

import "fmt"

func main() {

	fmt.Println("Enter a number: ")
	var input int

	// Get user input
	fmt.Scanf("%d", &input)

	// Create a slice of integers to store the prime numbers
	var primes []int

	// Create a slice of booleans to represent the prime numbers
	var isPrime []bool

	// Initialize the slice of booleans to true (all prime numbers)
	for i := 0; i < input; i++ {
		isPrime = append(isPrime, true)
	}

	// Loop through the slice of booleans and mark non-prime numbers as false
	for i := 2; i < input; i++ {
		if isPrime[i] {
			primes = append(primes, i)
			for j := i * i; j < input; j += i {
				isPrime[j] = false
			}
		}
	}

	// Print the prime numbers
	fmt.Println(primes)

}
