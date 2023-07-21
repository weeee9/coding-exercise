package pairSum

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// 使用快慢指標來快速找到中心
	slow := head
	fast := head
	stack := []int{}
	for fast != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	max := 0
	for slow != nil {
		val := slow.Val

		size := len(stack)
		pair := stack[size-1]
		sum := val + pair

		if sum > max {
			max = sum
		}

		stack = stack[:size-1]
		slow = slow.Next
	}

	return max
}
