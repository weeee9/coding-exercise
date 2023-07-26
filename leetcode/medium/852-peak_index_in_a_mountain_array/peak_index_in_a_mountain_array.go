package peakindexinamountainarray

func peakIndexInMountainArray(A []int) int {
	left, right := 0, len(A)

	// 使用二分搜索來尋找山峰的索引
	for left < right {
		mid := (left + right) / 2
		// 如果中間元素大於其左右相鄰元素，則該元素為山峰，返回其索引
		if A[mid-1] < A[mid] && A[mid] > A[mid+1] {
			return mid
		}

		// 如果中間元素大於其左相鄰元素，則山峰位於右半部分，更新左邊界為中間索引
		if A[mid] > A[mid-1] {
			left = mid
		}

		// 如果中間元素大於其右相鄰元素，則山峰位於左半部分，更新右邊界為中間索引
		if A[mid] > A[mid+1] {
			right = mid
		}
	}

	return 0
}
