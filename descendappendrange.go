package piscine

func DescendAppendRange(max, min int) []int {
	if max <= min {
		return []int{} // real empty slice, not nil
	}

	var result []int
	for i := max; i > min; i-- {
		result = append(result, i)
	}
	return result
}
