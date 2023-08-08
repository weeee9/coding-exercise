package mergeinbetweenlinkedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	merged, mergedHead := &ListNode{}, &ListNode{}

	count := 0
	curr1 := list1
	for curr1 != nil {
		if count < a || count > b {
			merged.Next = curr1
			if count == 0 {
				mergedHead = merged.Next
			}
			merged = merged.Next
			curr1 = curr1.Next
		} else if a <= count && count <= b {
			curr2 := list2
			if count == a {
				for curr2 != nil {
					merged.Next = curr2
					if count == 0 {
						mergedHead = merged.Next
					}
					curr2 = curr2.Next
					merged = merged.Next
				}
			}
			curr1 = curr1.Next
		}
		count++
	}

	return mergedHead
}
