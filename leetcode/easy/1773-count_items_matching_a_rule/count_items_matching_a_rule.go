package match

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	idx := 0
	switch ruleKey {
	case "type":
		idx = 0
	case "color":
		idx = 1
	case "name":
		idx = 2
	}

	match := 0
	for _, item := range items {
		if item[idx] == ruleValue {
			match++
		}
	}

	return match
}
