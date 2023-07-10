package sorting

/*
1. 首先，把數組分成已排序和未排序兩部分，初始時，已排序部分只有一個元素（即數組的第一個元素）。
2. 從未排序部分取出一個元素，和已排序部分的元素從後往前逐一比較，找到適合此元素的位置並插入。
3. 重複步驟2，直到所有元素都插入到已排序部分，排序完成。
*/
func InsertionSort(arr []int) []int {
	for currentIdx := 1; currentIdx < len(arr); currentIdx++ {
		key := arr[currentIdx]
		sortedPartStartIdx := currentIdx - 1

		// 將 key 之前的元素（已排序部分）向右移動，直到找到適合 key 的位置
		for sortedPartStartIdx >= 0 && arr[sortedPartStartIdx] > key {
			arr[sortedPartStartIdx+1] = arr[sortedPartStartIdx]

			sortedPartStartIdx--
		}

		// 插入 key
		arr[sortedPartStartIdx+1] = key
	}
	return arr
}
