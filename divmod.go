package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b // store the quotient in the variable pointed by div
	*mod = a % b // store the remainder in the variable pointed by mod
}
