package capacitytoshippackageswithinddays

func shipWithinDays(weights []int, days int) int {
	maxLoad, totalLoad := getMaxAndTotalLoad(weights)
	// 初始化左右邊界，左邊界為最大重量，右邊界為總重量
	left, right := maxLoad, totalLoad

	// 使用二分搜索來尋找最小船載重量
	for left < right {
		// 計算中間重量
		mid := left + (right-left)/2
		// 如果以中間重量為載重量，在指定天數內可以運送完所有貨物，則更新右邊界為中間重量
		if possible(weights, mid, days) {
			right = mid
		} else {
			// 否則更新左邊界為中間重量的下一個重量
			left = mid + 1
		}
	}

	// 最終返回左邊界，即為最小船載重量
	return left
}

// 檢查以指定載重量 cap，在指定天數 days 內是否可以運送完所有貨物
func possible(weights []int, cap, days int) bool {
	daysNeeded, currentLoad := 1, 0
	for _, weight := range weights {
		// 累加當前貨物的重量
		currentLoad += weight
		// 如果當前貨物的重量超過了載重量，則需要新增一天，將當前貨物的重量設置為新的貨物重量
		if currentLoad > cap {
			daysNeeded++
			currentLoad = weight
		}
	}

	// 檢查需要的天數是否小於等於指定的天數
	return daysNeeded <= days
}

func getMaxAndTotalLoad(weights []int) (max int, total int) {
	for i := range weights {
		total += weights[i]
		if weights[i] > max {
			max = weights[i]
		}
	}
	return
}
