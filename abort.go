package piscine

func Abort(a, b, c, d, e int) int {
	nums := []int{a, b, c, d, e}

	// Simple bubble sort since there are only 5 elements
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// Return the middle element (median)
	return nums[2]
}
