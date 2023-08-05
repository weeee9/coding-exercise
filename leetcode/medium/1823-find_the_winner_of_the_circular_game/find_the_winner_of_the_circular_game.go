package findthewinnerofthecirculargame

func findTheWinner(n int, k int) int {
	queue := make([]int, n)

	for i := 0; i < n; i++ {
		queue[i] = i + 1
	}

	for len(queue) > 1 {
		for count := 1; count <= k; count++ {
			pop := queue[0]
			queue = queue[1:]

			// it's not the target, push it back
			if count != k {
				//
				queue = append(queue, pop)
			}
		}
	}

	return queue[0]
}
