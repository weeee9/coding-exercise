package add

func minAddToMakeValid(s string) int {
	stack := []byte{}

	for i := range s {
		lenOfStack := len(stack)
		lenOfStackGreaterThanZero := lenOfStack > 0

		if !lenOfStackGreaterThanZero {
			stack = append(stack, s[i])
			continue
		}

		isRightParentheses := s[i] == ')'
		lastElementIsLeftParentheses := stack[lenOfStack-1] == '('

		if isRightParentheses && lastElementIsLeftParentheses {
			stack = stack[:lenOfStack-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack)
}
