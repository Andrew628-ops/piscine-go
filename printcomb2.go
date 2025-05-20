package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	// Loop through the first number 'a'
	for a := 0; a <= 98; a++ {
		// Loop through the second number 'b', which should be greater than 'a'
		for b := a + 1; b <= 99; b++ {
			// Print first number 'a'
			PrintDigit(a / 10) // Tens digit of 'a'
			PrintDigit(a % 10) // Ones digit of 'a'
			z01.PrintRune(' ') // Space between the two numbers

			// Print second number 'b'
			PrintDigit(b / 10) // Tens digit of 'b'
			PrintDigit(b % 10) // Ones digit of 'b'

			// Print a comma and space, except for the last combination (98 99)
			if !(a == 98 && b == 99) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	// Print a newline after printing all combinations
	z01.PrintRune('\n')
}

// Helper function to print a single digit
func PrintDigit(n int) {
	digits := [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	z01.PrintRune(digits[n]) // Print the corresponding rune (digit)
}
