/*
fibonacci.go
An example of generating the fibonacci sequence.

The fibonacci sequence is a series of numbers, in which the current number 
in the sequence is the sum of adding the two previous numbers.

Author: Ricky Hewitt <ricky@rickyhewitt.dev>
*/

package main

import "fmt"

func main() {

	// Define limit
	// Fibonacci sequence can grow quite large... so we use uint64 type throughout this program.
	const LIMIT uint64 = 50

	// Initialize an array, which will store ALL numbers.
	// Note: It would be more efficient if I only stored the two prior numbers,
	// 		 but this will suffice for this example.
	var numbers []uint64

	// Use a string var for output
	// This is to avoid multiple unnecessary calls to fmt.Printf (expensive).
	var output string = ""

	// Begin loop
	for i := uint64(0); i < LIMIT; i++ {

		// Construct current number.		
		var currentNum uint64

		// Current number is the sum of the two previous numbers added together.
		if len(numbers) > 1 {
			currentNum = numbers[i-1] + numbers[i-2]
		} else {
			// Initial starting number on first initiation (0 by default).
			currentNum = i
		}

		// Append (push) current number to Stack
		numbers = append(numbers, currentNum)

		// Append current number to output
		output = fmt.Sprintf("%s\n%v", output, currentNum)

	}

	// Display final output
	fmt.Println(output)
}
