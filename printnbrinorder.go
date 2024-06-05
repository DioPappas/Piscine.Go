package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return // do nothing for negative numbers
	}

	// Find the number of digits in the integer
	numDigits := countDigits(n)

	// Extract and print the digits one by one
	for i := 0; i < numDigits; i++ {
		smallestDigit := findSmallestDigit(n)
		if smallestDigit == 0 && n != 0 {
			z01.PrintRune('0') // Print the digit 0
		} else {
			z01.PrintRune(rune(smallestDigit) + '0') // Print the digit
		}
		n = removeDigit(n, smallestDigit) // Remove the printed digit from the integer
	}

	// Print a newline character at the end
	z01.PrintRune('\n')
}

// Function to count the number of digits in an integer
func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

// Function to find the smallest digit in an integer
func findSmallestDigit(n int) int {
	smallest := 10 // Initialize with a value greater than any digit
	for n != 0 {
		digit := n % 10
		if digit < smallest {
			smallest = digit
		}
		n /= 10
	}
	return smallest
}

// Function to remove a specific digit from an integer
func removeDigit(n, digit int) int {
	result := 0
	tens := 1
	for n != 0 {
		currDigit := n % 10
		if currDigit != digit {
			result += currDigit * tens
			tens *= 10
		}
		n /= 10
	}
	return result
}
