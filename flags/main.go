package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	printStr("--insert")
	newline()
	printStr("  -i")
	newline()
	printStr("         This flag inserts the string into the string passed as argument.")
	newline()
	printStr("--order")
	newline()
	printStr("  -o")
	newline()
	printStr("         This flag will behave like a boolean, if it is called it will order the argument.")
	newline()
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}

func newline() {
	z01.PrintRune('\n')
}

func sortRunes(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-1-i; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}
	return string(runes)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	inputStr := ""
	insertStr := ""
	order := false

	for _, arg := range args {
		if len(arg) >= 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) >= 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else {
			inputStr = arg
		}
	}

	result := inputStr + insertStr

	if order {
		result = sortRunes(result)
	}

	printStr(result)
	newline()
}
