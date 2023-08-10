package reversesubstringsbetweeneachpairofparentheses

func reverseParentheses(s string) string {
	stack := make([]int, 0, len(s))
	result := ""
	bracketCount := 0
	for idx, r := range s {
		if r == '(' {
			realLeftIdx := idx - bracketCount
			stack = append(stack, realLeftIdx)
			bracketCount++
		} else if r == ')' {
			realLeftIdx := stack[len(stack)-1]
			realRightIdx := idx - bracketCount - 1

			result = reverse([]rune(result), realLeftIdx, realRightIdx)

			stack = stack[:len(stack)-1]
			bracketCount++
		} else {
			result += string(r)
		}
	}

	return result
}

func reverse(s []rune, left, right int) string {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return string(s)
}
