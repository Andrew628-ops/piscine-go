package piscine

func PodiumPosition(podium [][]string) [][]string {
	var reversed [][]string

	for i := len(podium) - 1; i >= 0; i-- {
		reversed = append(reversed, podium[i])
	}

	return reversed
}
