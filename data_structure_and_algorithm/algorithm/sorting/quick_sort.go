package sorting

func quickSort(array []int) []int {
	// 如果數組長度小於2，則返回數組本身（因為它已經排序了）
	if len(array) < 2 {
		return array
	} else {
		// 選擇一個基準點（此處我們選擇了數組的第一個元素）
		pivot := array[0]
		// less 和 greater 分別用於存放比基準點小的元素和大的元素
		var less []int
		var greater []int
		// 對數組中的每一個元素進行分區
		for _, num := range array[1:] {
			if num <= pivot {
				less = append(less, num) // 比基準點小的元素放入 less
			} else {
				greater = append(greater, num) // 比基準點大的元素放入 greater
			}
		}
		// 使用遞歸對 less 和 greater 進行排序，並將結果合併
		var result []int
		result = append(result, quickSort(less)...)
		result = append(result, pivot)
		result = append(result, quickSort(greater)...)
		return result
	}
}
