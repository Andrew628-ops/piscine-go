package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	aStr := os.Args[1]
	op := os.Args[2]
	bStr := os.Args[3]

	a, ok1 := Atoi(aStr)
	b, ok2 := Atoi(bStr)

	if !ok1 || !ok2 {
		return
	}

	var result int
	var output string

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			output = "No division by 0\n"
			os.Stdout.Write([]byte(output))
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			output = "No modulo by 0\n"
			os.Stdout.Write([]byte(output))
			return
		}
		result = a % b
	default:
		return
	}

	output = Itoa(result) + "\n"
	os.Stdout.Write([]byte(output))
}
