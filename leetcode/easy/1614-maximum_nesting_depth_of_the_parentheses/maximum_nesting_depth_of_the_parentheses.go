package max

func maxDepth(s string) int {
	leftParentheses := []byte{}

	maxDepth := 0
	for i := range s {
		if s[i] == '(' {
			leftParentheses = append(leftParentheses, s[i])
		}

		if s[i] == ')' {
			depth := len(leftParentheses)
			if depth > maxDepth {
				maxDepth = depth
			}
			leftParentheses = leftParentheses[:depth-1]
		}
	}

	return maxDepth
}
