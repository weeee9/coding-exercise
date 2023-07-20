package construct

import "fmt"

func smallestNumber(pattern string) string {
	result := ""
	stack := []int{}
	for i := 0; i <= len(pattern); i++ {
		stack = append(stack, i+1)
		if i == len(pattern) || pattern[i] == 'I' {
			for len(stack) > 0 {
				size := len(stack)
				result += fmt.Sprint(stack[size-1])
				stack = stack[:size-1]
			}
		}
	}

	return result
}
