/*
   pi-to-nth.go

   Example program for outputting pi, with a limit
   supplied by the user.

   Note: This example does not generate pi itself, it uses the math.Pi constant provided by the
		 math package & formats accordingly.

   ricky@rickyhewitt.dev
*/

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var input string
	var limit int

	// Initial program output
	fmt.Println("pi-to-nth-digit")
	fmt.Println("How many digits of pi would you like?")

	// Use fmt.Scan to get keyboard input
	fmt.Print("Digits: ")
	_, err := fmt.Scan(&input)

	// Check for errors
	if err == nil {
		conv, err2 := strconv.Atoi(input)
		if err2 == nil {
			// Check for valid range
			if (conv > 0) && (conv <= 48) {
				// Valid range, set limit
				limit = conv
			} else {
				fmt.Println("Error: Invalid range. Valid length is 1 - 48 digits.")
				os.Exit(1)
			}
		} else {
			fmt.Printf("Error: '%s' is not a valid number.\n", input)
			os.Exit(1)
		}
	} else {
		fmt.Printf("Error: Invalid input.\n")
		os.Exit(1)
	}

	// Format the precision to the requested decimal places
	format := fmt.Sprintf("Result: %%.%df\n", limit)

	// Print the formatted value of math.Pi
	fmt.Printf(format, math.Pi)

}
