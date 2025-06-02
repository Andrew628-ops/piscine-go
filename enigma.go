package piscine

func Enigma(a ***int, b *int, c *******int, d ****int) {
	// Store the values
	valA := ***a
	valC := *******c
	valD := ****d
	valB := *b

	// Move the values in circular order
	*******c = valA
	****d = valC
	*b = valD
	***a = valB
}
