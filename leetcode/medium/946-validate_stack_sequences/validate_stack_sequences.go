package validatestacksequences

func validateStackSequences(pushed []int, popped []int) bool {
	// create a stack to do the actual push and pop action
	stack := make([]int, 0, len(pushed))
	poppedIdx := 0

	for _, val := range pushed {
		stack = append(stack, val)

		// do pop action if the top of stack is 
		// same as the popped[idx] element
		for len(stack) > 0 && stack[len(stack)-1] == popped[poppedIdx] {
			stack = stack[:len(stack)-1]
			poppedIdx++
		}
	}

	// if stack sequence is valid, we should see
	// every pushed element was popped from stack
	return len(stack) == 0
}
