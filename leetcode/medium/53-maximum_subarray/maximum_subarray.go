package maximumsubarray

func maxSubArray(nums []int) int {
	// 初始化當前和與最大和為陣列的第一個元素
	current, max := nums[0], nums[0]

	// 遍歷陣列的其他元素，找到和最大的子陣列的和
	for _, num := range nums[1:] {
		// 如果當前和小於 0，將當前和設置為當前元素值，從當前元素開始重新計算子陣列和
		if current < 0 {
			current = num
		} else {
			// 否則將當前元素值加到當前和上，擴展子陣列的範圍
			current += num
		}

		if current > max {
			max = current
		}
	}

	return max
}
