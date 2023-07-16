package valid

func isValid(s string) bool {
	openParens := map[rune]int{'(': 1, '{': 2, '[': 3}
	closingParens := map[rune]int{')': 1, '}': 2, ']': 3}

	// 使用 stack 來處理括號配對
	stack := []rune{}
	for _, bracket := range s {
		// 如果是開放括號，將其推入 stack 中
		if _, ok := openParens[bracket]; ok {
			stack = append(stack, bracket)
			continue
		}

		// 如果是閉合括號，檢查 stack 中最新的開放括號是否和當前閉合括號匹配
		if _, ok := closingParens[bracket]; ok {
			if len(stack) <= 0 {
				return false
			}

			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if openParens[pop] != closingParens[bracket] {
				return false
			}
		}
	}

	// 如果 stack 還有剩餘的括號，表示有未配對的括號，括號無效
	return len(stack) == 0
}
