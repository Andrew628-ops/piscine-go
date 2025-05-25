package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func printRune(r rune) {
	z01.PrintRune(r)
}

func printSpace() {
	z01.PrintRune(' ')
}

func printNewline() {
	z01.PrintRune('\n')
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printNewline()
		return
	}

	var vowels []rune
	for _, word := range args {
		for _, ch := range word {
			if isVowel(ch) {
				vowels = append(vowels, ch)
			}
		}
	}

	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	vowelIndex := 0
	for i, word := range args {
		for _, ch := range word {
			if isVowel(ch) && vowelIndex < len(vowels) {
				printRune(vowels[vowelIndex])
				vowelIndex++
			} else {
				printRune(ch)
			}
		}
		if i != len(args)-1 {
			printSpace()
		}
	}
	printNewline()
}
