package buildarray

func buildArray(target []int, n int) []string {
	result := make([]string, 0)

	idxOfTarget := 0
	for num := 1; num <= n; num++ {
		if idxOfTarget == len(target) {
			break
		}

		result = append(result, "Push")

		// num 並不存在於 target 中，pop 
		if num != target[idxOfTarget] {
			result = append(result, "Pop")
		} else {
			idxOfTarget++
		}
	}

	return result
}
