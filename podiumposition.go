package piscine

func PodiumPosition(podium [][]string) [][]string {
	reversed := make([][]string, len(podium))

	for i := 0; i < len(podium); i++ {
		reversed[i] = podium[len(podium)-1-i]
	}

	return reversed
}
