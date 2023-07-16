package sorting

// 實現合併排序
func mergeSort(array []int) []int {
	// 基本情況：如果 array 長度小於2，則返回 array 本身（因為它已經排序了）
	if len(array) < 2 {
		return array
	}

	// 分治：將數組分為兩部分
	mid := len(array) / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	// 合併：將兩部分已排序的 array 合併為一個有序的 array
	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

	// 將左右兩側的元素逐一進行比較，並將較小的元素添加到結果中
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	// 將剩餘的元素添加到結果中
	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
