package search

func search(nums []int, target int) int {
	pivot := findPivot(nums)

	var left, right int
	switch {
	case pivot == -1:
		left, right = 0, len(nums)-1

	// pivot = 4, nums = [5, 6, 7, 8, 0, 1, 2], target = 2
	// only needs to search from [pivot to last idx]
	case target < nums[0]:
		left, right = pivot, len(nums)-1

	// pivot = 4, nums = [5, 6, 7, 8, 0, 1, 2], target = 8
	// pivot = 1, nums = [3, 1]
	// only needs to search from [first idx to pivot]
	default:
		left, right = 0, pivot-1
	}

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}

	if nums[left] == target {
		return left
	}

	return -1
}

func findPivot(nums []int) int {
	if len(nums) == 1 {
		return -1
	}
	left, right := 0, len(nums)-1

	if nums[left] < nums[right] {
		return -1
	}

	for left < right {
		mid := (left + right) / 2

		switch {
		//  l    m	     r
		// [..., 7, 0, ...]
		case nums[mid] > nums[mid+1]:
			return mid + 1

		//  l       m	 r
		// [..., 7, 0, ...]
		case nums[mid] < nums[mid-1]:
			return mid

		//  l       m	 r
		// [1, ..., 7, ...]
		case nums[left] < nums[mid]:
			left = mid + 1

		//  l       m	 r
		// [1, ..., 7, ...]
		default:
			right = mid - 1
		}
	}

	return -1
}
