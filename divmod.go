package piscine

func DivMod(a init, b init, div *init, mod *init) {
	*div = a / b // store the quotient in the variable pointed by div
	*mod = a % b // store the remainder in the variable pointed by mod
}
