package optimalpartitionofstring

func partitionString(s string) int {
	count := 1
	exist := map[rune]struct{}{}

	for _, c := range s {
		if _, ok := exist[c]; ok {
			count++
			exist = map[rune]struct{}{}
		}
		exist[c] = struct{}{}
	}

	return count
}
