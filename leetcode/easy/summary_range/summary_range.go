package summaryrange

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	smallestSortedList := []string{}
	headValue := nums[0]
	for idx := 0; idx < len(nums); idx++ {
		if idx < len(nums)-1 &&
			nums[idx]+1 == nums[idx+1] {
			continue
		}

		if headValue == nums[idx] {
			smallestSortedList = append(smallestSortedList, fmt.Sprint(headValue))
		} else {
			smallestSortedList = append(smallestSortedList, fmt.Sprintf("%d->%d", headValue, nums[idx]))
		}

		if idx < len(nums)-1 {
			headValue = nums[idx+1]
		}
	}

	return smallestSortedList
}
