package build

func buildArray(nums []int) []int {
	result := make([]int, len(nums))

	for i := range nums {
		result[i] = nums[nums[i]]
	}

	return result
}

func buildArrayWithO1Space(nums []int) []int {
	for i := range nums {
		nums[i] += (nums[nums[i]] % 1000) * 1000
	}

	for i := range nums {
		nums[i] = nums[i] / 1000
	}

	return nums
}
