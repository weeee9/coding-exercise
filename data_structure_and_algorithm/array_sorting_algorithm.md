# Array Sorting Algorithm

![image](https://i0.wp.com/miro.medium.com/max/596/1*ipkeWQ_Lb0lbkhB8rigxTA.png?w=1170&ssl=1)

## Insertion Sort


插入排序（Insertion Sort）是一種簡單且直觀的排序算法。其工作原理類似於我們玩撲克牌，將每一張牌插入到其他已經有序的牌中的適當位置。

![image](https://i0.wp.com/turingplanet.org/wp-content/uploads/2020/02/InsertionSortGIF.gif?resize=523%2C325&ssl=1)

以下是插入排序的基本步驟：


1. 首先，把數組分成已排序和未排序兩部分，初始時，已排序部分只有一個元素（即數組的第一個元素）。
2. 從未排序部分取出一個元素，和已排序部分的元素從後往前逐一比較，找到適合此元素的位置並插入。
3. 重複步驟2，直到所有元素都插入到已排序部分，排序完成。

在最壞的情況下（當輸入數據是反序排列時），插入排序的時間複雜度為 O(n²)。然而在最好的情況下（當輸入數據已經排序時），其時間複雜度為 O(n)。這使得插入排序在數據近似有序的情況下表現良好。

插入排序在實際應用中的一個主要優點是，它只需要固定的額外空間 O(1)，因此對於空間敏感的情況，插入排序是一種實用的選擇。另外，對於小數據集，插入排序也是一種效率高的選擇，因為其常數因子相對較小。

```go
func insertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		key := array[i]
		j := i - 1

		// 將 key 之前的元素（已排序部分）向右移動，直到找到適合 key 的位置
		for j >= 0 && array[j] > key {
			array[j+1] = array[j]
			j = j - 1
		}
		// 插入 key
		array[j+1] = key
	}
	return array
}
```


## Quick Sort


快速排序（Quick Sort）是一種使用分治策略(Divide and Conquer)的高效排序算法。

![image](https://i0.wp.com/images.deepai.org/glossary-terms/a5228ea07c794b468efd1b7f758b9ead/Quicksort.png?resize=577%2C410&ssl=1)

快速排序的基本步驟如下：


1. 選擇數據中的一個元素作為「基準點」（pivot）。
2. 將數據分為兩部分，一部分是比基準點小的元素，另一部分是比基準點大的元素。這個操作被稱為分區（partitioning）。
3. 對兩部分數據分別進行快速排序。

重複以上步驟，直到所有數據都已排序。

在最壞的情況下（輸入數據已排序或反序排列），快速排序的時間複雜度為 O(n²)。然而在平均情況下，其時間複雜度為 O(n log n)。實際上，由於其良好的緩存性能和內部迭代方式，對於大型數據集，快速排序通常優於其他 O(n log n) 的排序算法。

```go
func quickSort(array []int) []int {
	// 如果 array 長度小於2，則返回 array 本身（因為它已經排序了）
	if len(array) < 2 {
		return array
	} else {
		// 選擇一個基準點（此處我們選擇了 array 數組的第一個元素）
		pivot := array[0]

		// less 和 greater 分別用於存放比基準點小的元素和大的元素
		var less []int
		var greater []int

		// 對 array 中的每一個元素進行分區
		for _, num := range array[1:] {
			if num <= pivot {
				less = append(less, num)  // 比基準點小的元素放入 less
			} else {
				greater = append(greater, num)  // 比基準點大的元素放入 greater
			}
		}
		// 使用遞迴對 less 和 greater 進行排序，並將結果合併
		var result []int
		result = append(result, quickSort(less)...)
		result = append(result, pivot)
		result = append(result, quickSort(greater)...)
		return result
	}
}
```


## Merge Sort


合併排序（Merge Sort）也是一種使用分治策略的有效排序算法。

![image](https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/Merge_sort_algorithm_diagram.svg/618px-Merge_sort_algorithm_diagram.svg.png)

合併排序的基本步驟如下：


1. 將數據分為兩部分。
2. 分別對兩部分數據進行合併排序。
3. 將兩部分已排序的數據合併為一個有序的數據。

重複以上步驟，直到所有數據都已排序。

無論在最好、最壞還是平均情況下，合併排序的時間複雜度都是 O(n log n)。然而，由於合併操作需要額外的空間，合併排序的空間複雜度是 O(n)。

```go
// 分治並合併兩個有序 slice
func merge(left, right []int) []int {
	var result []int
	i, j := 0, 0

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
```