package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if table[j] > table[j+1] {
				// Swap values
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
