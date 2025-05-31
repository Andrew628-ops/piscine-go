package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Check for exactly 3 arguments (program name + 3 args)
	if len(os.Args) != 4 {
		return
	}

	// Parse the first number
	a, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		return
	}

	// Parse the operator
	op := os.Args[2]
	if len(op) != 1 || !isValidOperator(op[0]) {
		return
	}

	// Parse the second number
	b, err := strconv.ParseInt(os.Args[3], 10, 64)
	if err != nil {
		return
	}

	// Perform the operation with overflow checks
	var result int64
	var overflow bool

	switch op {
	case "+":
		result, overflow = addWithOverflow(a, b)
	case "-":
		result, overflow = subtractWithOverflow(a, b)
	case "*":
		result, overflow = multiplyWithOverflow(a, b)
	case "/":
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}
		result = a % b
	}

	if overflow {
		return
	}

	fmt.Println(result)
}

func isValidOperator(op byte) bool {
	return op == '+' || op == '-' || op == '*' || op == '/' || op == '%'
}

func addWithOverflow(a, b int64) (int64, bool) {
	if b > 0 && a > (1<<63-1)-b {
		return 0, true
	}
	if b < 0 && a < (-1<<63)-b {
		return 0, true
	}
	return a + b, false
}

func subtractWithOverflow(a, b int64) (int64, bool) {
	if b < 0 && a > (1<<63-1)+b {
		return 0, true
	}
	if b > 0 && a < (-1<<63)+b {
		return 0, true
	}
	return a - b, false
}

func multiplyWithOverflow(a, b int64) (int64, bool) {
	if a == 0 || b == 0 {
		return 0, false
	}

	result := a * b
	if a == -1<<63 && b == -1 {
		return 0, true // Special case: minimum int64 * -1 overflows
	}
	if result/b != a {
		return 0, true
	}
	return result, false
}
