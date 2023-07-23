package pascalstriangle

func generate(numRows int) [][]int {
	pascalTriangle := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		if i == 0 {
			pascalTriangle[i] = []int{1}
		} else {
			pascalTriangle[i] = generateNextRow(pascalTriangle[i-1])
		}
	}

	return pascalTriangle
}

func generateNextRow(prevRow []int) []int {
	nextRow := make([]int, len(prevRow)+1)

	for i := 0; i < len(nextRow); i++ {
		if i == 0 || i == len(nextRow)-1 {
			nextRow[i] = 1
		} else {
			nextRow[i] = prevRow[i-1] + prevRow[i]
		}
	}

	return nextRow
}
