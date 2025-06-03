package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := len(podium)
	var result [][]string

	// Create a slice of correct length by assigning manually
	result = podium

	for i := 0; i < length/2; i++ {
		// Swap elements
		podium[i], podium[length-1-i] = podium[length-1-i], podium[i]
	}
	return result
}
