package main

import "os"

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
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
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

func Atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}

	sign := 1
	i := 0
	if s[0] == '-' {
		sign = -1
		i++
	} else if s[0] == '+' {
		i++
	}

	num := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		num = num*10 + int(s[i]-'0')
	}
	return num * sign, true
}

func Itoa(n int) string {
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
