package move

func minOperations(boxes string) []int {
	result := make([]int, len(boxes))

	var balls, lastDis int
	for i := 1; i < len(boxes); i++ {
		if boxes[i-1] == '1' {
			balls++
		}
		lastDis += balls
		result[i] += lastDis
	}

	balls, lastDis = 0, 0
	for i := len(boxes) - 2; i >= 0; i-- {
		if boxes[i+1] == '1' {
			balls++
		}
		lastDis += balls
		result[i] += lastDis
	}

	return result
}
