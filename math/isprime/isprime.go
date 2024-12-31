/*
   isprime.go
   An example of checking if a number is prime (golang).

   Author: Ricky Hewitt <ricky@rickyhewitt.dev>
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// isPrime returns true or false depending on whether or not a number is prime
func isPrime(n int) bool {

	// Number <= 1 cannot be prime
	if n <= 1 {
		return false
	}
	// Test division up to the sqrt of n
	// math.Sqrt expects a float64, so we have to cast int -> float64.
	for i := 2; float64(i) <= math.Sqrt(float64(n)); i++ {
		if n%i == 0 {
			return false
		}
	}

	// Otherwise prime
	return true

}

func main() {
	var input string

	// Initial program output
	fmt.Println("isprime.go")
	fmt.Println("Which number would you like to check?")

	// Use fmt.Scan to get keyboard input
	fmt.Print("Number: ")
	_, err := fmt.Scan(&input)

	// Check for errors
	if err != nil {
		fmt.Printf("Error: Invalid input.\n")
		os.Exit(1)
	} else {
		// Convert input string to integer
		conv, err2 := strconv.Atoi(input)

		// Error handling
		if err2 != nil {
			fmt.Printf("Error: '%s' is not a valid number.\n", input)
			os.Exit(1)
		} else {
			// Valid type, can now check for prime
			prime := isPrime(conv)
			if prime {
				fmt.Printf("%d is a prime number\n", conv)
			} else {
				fmt.Printf("%d is NOT a prime number\n", conv)
			}

		}
	}
}
