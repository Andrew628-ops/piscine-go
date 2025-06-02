package piscine

func Compact(ptr *[]string) int {
	original := *ptr
	compacted := []string{}

	for _, str := range original {
		if str != "" {
			compacted = append(compacted, str)
		}
	}
	*ptr = compacted
	return len(compacted)
}
