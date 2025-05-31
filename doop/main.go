package main

import "os"

const MaxInt64 = 9223372036854775807
const MinInt64 = -9223372036854775808

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

	var result int64
	switch op {
	case "+":
		if (b > 0 && a > MaxInt64-b) || (b < 0 && a < MinInt64-b) {
			return
		}
		result = a + b
	case "-":
		if (b < 0 && a > MaxInt64+b) || (b > 0 && a < MinInt64+b) {
			return
		}
		result = a - b
	case "*":
		if (a > 0 && b > 0 && a > MaxInt64/b) ||
			(a > 0 && b < 0 && b < MinInt64/a) ||
			(a < 0 && b > 0 && a < MinInt64/b) ||
			(a < 0 && b < 0 && a < MaxInt64/b) {
			return
		}
		result = a * b
	case "/":
		if b == 0 {
			os.Stdout.Write([]byte("No division by 0\n"))
			return
		}
		result = a / b
	case "%":
		if b == 0 {
			os.Stdout.Write([]byte("No modulo by 0\n"))
			return
		}
		result = a % b
	default:
		return
	}

	os.Stdout.Write([]byte(Itoa(result) + "\n"))
}

func Atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := int64(1)
	i := 0
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	var num int64 = 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int64(s[i] - '0')
		if num > (MaxInt64-digit)/10 {
			return 0, false // overflow during parsing
		}
		num = num*10 + digit
	}
	return num * sign, true
}

func Itoa(n int64) string {
	if n == 0 {
		return "0"
	}

	sign := ""
	if n < 0 {
		sign = "-"
		n = -n
	}

	var digits []byte
	for n > 0 {
		digits = append([]byte{byte(n%10) + '0'}, digits...)
		n /= 10
	}

	return sign + string(digits)
}
