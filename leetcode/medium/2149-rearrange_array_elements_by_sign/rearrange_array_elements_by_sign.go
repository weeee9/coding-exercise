package rearrangearrayelementsbysign

func rearrangeArray(nums []int) []int {
	// create an array to store the result
	result := make([]int, len(nums))
	// create two pointers, one for positive numbers and one for negative numbers
	positiveIdx := 0
	negativeIdx := 1

	for _, num := range nums {
		// if the number is positive, store it at the positive pointer
		if num > 0 {
			result[positiveIdx] = num
			// move the positive pointer 2 spaces to the right
			positiveIdx += 2
		} else {
			// if the number is negative, store it at the negative pointer
			result[negativeIdx] = num
			// move the negative pointer 2 spaces to the right
			negativeIdx += 2
		}
	}

	return result
}

func rearrangeArrayUsingRotate(nums []int) []int {
	targetIdx := -1
	for i := range nums {
		switch {
		case targetIdx >= 0:

		case targetIdx == -1:
			numIsPositive := nums[i] >= 0
			idxIsEven := i%2 == 0

			positiveNumAtEvenIdx := numIsPositive && idxIsEven
			negativeNumAtOddIdx := !numIsPositive && !idxIsEven

			if !positiveNumAtEvenIdx || !negativeNumAtOddIdx {
				targetIdx = i
			}
		}
	}

	return nums
}

