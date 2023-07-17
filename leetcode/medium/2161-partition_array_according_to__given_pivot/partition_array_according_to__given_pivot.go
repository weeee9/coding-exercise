package partition

func pivotArray(nums []int, pivot int) []int {
	result := make([]int, len(nums))
	idx := 0

	for _, num := range nums {
		if num < pivot {
			result[idx] = num
			idx++
		}
	}

	for _, num := range nums {
		if num == pivot {
			result[idx] = num
			idx++
		}
	}

	for _, num := range nums {
		if num > pivot {
			result[idx] = num
			idx++
		}
	}

	return result
}
