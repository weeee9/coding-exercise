package dailytemperatures

func dailyTemperatures(temperatures []int) []int {
	days := make([]int, len(temperatures))
	stack := make([]int, 0, len(temperatures))

	for currentIdx, currentTemperature := range temperatures {
		// 檢查堆疊是否有元素，並且當前溫度是否高於堆疊頂部索引對應的溫度
		for len(stack) > 0 && currentTemperature > temperatures[stack[len(stack)-1]] {
			// 彈出堆疊頂部索引
			lastIdx := stack[len(stack)-1]

			// 計算當前溫度與彈出的索引對應的溫度之間的天數差距
			distance := currentIdx - lastIdx
			// 將該天數差距存儲到結果陣列中
			days[lastIdx] = distance

			// 將該索引從堆疊中移除
			stack = stack[:len(stack)-1]
		}

		// 將當前溫度的索引推入堆疊中，供未來的比較使用
		stack = append(stack, currentIdx)
	}

	return days
}
