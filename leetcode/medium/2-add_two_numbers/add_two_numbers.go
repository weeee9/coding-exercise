package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addIterative(l1, l2)
}

func addIterative(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	numA, numB := l1, l2
	carry := 0

	for numA != nil || numB != nil {
		sum := numA.Val + numB.Val + carry
		carry = sum / 10

		current.Next = &ListNode{Val: sum % 10}
		current = current.Next

		if numA.Next != nil || numB.Next != nil {
			if numA.Next == nil {
				numA.Next = &ListNode{Val: 0}
			}
			if numB.Next == nil {
				numB.Next = &ListNode{Val: 0}
			}
		} else if carry > 0 {
			current.Next = &ListNode{Val: carry}
			return result.Next
		}

		numA = numA.Next
		numB = numB.Next
	}

	return result.Next
}
