package pascalstriangle

func generate(numRows int) [][]int {
	pascalTriangle := make([][]int, 0, numRows)

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else if i > 1 {
				a := pascalTriangle[i-1][j-1]
				b := pascalTriangle[i-1][j]
				row[j] = a + b
			}
		}
		pascalTriangle = append(pascalTriangle, row)
	}

	return pascalTriangle
}
