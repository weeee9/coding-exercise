package queriesonnumberofpointsinsideacircle

func countPoints(points [][]int, queries [][]int) []int {
	result := make([]int, len(queries))

	for i, query := range queries {
		for _, point := range points {
			if pointIsCircle(query, point) {
				result[i]++
			}
		}
	}

	return result
}

func pointIsCircle(circle, point []int) bool {
	px, py := point[0], point[1]
	cx, cy := circle[0], circle[1]
	r := circle[2]

	dx, dy := px-cx, py-cy

	return (dx*dx)+(dy*dy) <= r*r
}
