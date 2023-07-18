package wildest

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	result := 0
	for i := 0; i < len(points)-1; i++ {
		result = max(result, points[i+1][0]-points[i][0])
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
