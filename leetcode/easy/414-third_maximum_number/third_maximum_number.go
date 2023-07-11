package thirdmaximumnumber

import "math"

func thirdMax(nums []int) int {
	first, second, third := math.MinInt64, math.MinInt64, math.MinInt64
	for i := 0; i < len(nums); i++ {
		if nums[i] == first || nums[i] == second || nums[i] == third {
			continue
		}

		if nums[i] > first {
			first, second, third = nums[i], first, second
		} else if nums[i] > second {
			second, third = nums[i], second
		} else if nums[i] > third {
			third = nums[i]
		}
	}

	if third == math.MinInt64 {
		return first
	}

	return third
}
