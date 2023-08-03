package findplayerswithzerooronelosses

import "sort"

func findWinners(matches [][]int) [][]int {
	winnings := map[int]int{}
	losses := map[int]int{}

	for _, match := range matches {
		winner := match[0]
		loser := match[1]

		winnings[winner]++

		losses[loser]++
	}

	zero := make([]int, 0)
	for player := range winnings {
		_, hasLosses := losses[player]
		if !hasLosses {
			zero = append(zero, player)
		}
	}

	ones := make([]int, 0)
	for player, count := range losses {
		if count == 1 {
			ones = append(ones, player)
		}
	}

	sort.Ints(zero)
	sort.Ints(ones)

	return [][]int{zero, ones}
}
