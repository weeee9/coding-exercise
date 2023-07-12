package skyline

func maxIncreaseKeepingSkyline(grid [][]int) int {
	n := len(grid)
	// 紀錄每一個 col 的最大值
	northToSouthSkyline := make([]int, n)
	// 紀錄每一個 row 的最大值
	westToEastSkyline := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			northToSouthSkyline[j] = max(northToSouthSkyline[j], grid[i][j])
			westToEastSkyline[i] = max(westToEastSkyline[i], grid[i][j])
		}
	}

	sum := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 增加的高度不能超過該位置對應的 row 和 col 的最大值
			// 所以要取對應 row 和 col 最大值中較小的那個值
			// 如: row_1 的最大值為 5, col_1 的最大值為 7
			// 為了不破壞 skyline, 所以只能增加到 5 而不是 7
			sum += min(northToSouthSkyline[j], westToEastSkyline[i]) - grid[i][j]
		}
	}

	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
