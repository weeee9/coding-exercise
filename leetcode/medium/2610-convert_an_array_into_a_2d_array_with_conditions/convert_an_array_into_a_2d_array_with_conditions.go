package convert

func findMatrix(nums []int) [][]int {
	numberCounter := map[int]int{}
	maxCount := 0
	for i := range nums {
		numberCounter[nums[i]]++
		if numberCounter[nums[i]] > maxCount {
			maxCount = numberCounter[nums[i]]
		}
	}

	result := make([][]int, maxCount)

	for num, count := range numberCounter {
		for i := 0; i < count; i++ {
			result[i] = append(result[i], num)
		}
	}

	return result
}

func anotherSolution(nums []int) [][]int {
	result := make([][]int, 0)

	idx := 0
	put := map[int]struct{}{}
	row := make([]int, 0)
	for len(put) < len(nums) {
		if idx == len(nums) {
			result = append(result, row)
			idx = 0
			row = make([]int, 0)
		}

		if _, hasPut := put[idx]; hasPut {
			idx++
			continue
		}

		if !in(nums[idx], row) {
			row = append(row, nums[idx])
			put[idx] = struct{}{}
		}

		idx++
		if len(put) == len(nums) {
			result = append(result, row)
			idx = 0
			row = make([]int, 0)
		}
	}

	return result
}

func in(n int, arr []int) bool {
	for i := range arr {
		if arr[i] == n {
			return true
		}
	}

	return false
}
