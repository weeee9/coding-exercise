package mergenodesinbetweenzeros

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	current := head.Next

	var mergedHead, temp *ListNode
	sum := 0

	for current != nil {
		// 如果當前節點的值為 0，則將連續的數字 0 合併成單一節點
		if current.Val == 0 {
			// 如果 mergedHead 為 nil，就建立一個新的 Node
			if mergedHead == nil {
				mergedHead = &ListNode{
					Val: sum,
				}

				// 接下來都使用 temp 操作，所以也要將 temp 指向 mergedHead
				temp = mergedHead
			} else {
				// 將中間非 0 的數字總和當作下一個節點的值
				temp.Next = &ListNode{
					Val: sum,
				}

				temp = temp.Next
			}

			// 重置 sum 為 0，用於計算下一個連續數字 0 之後的數字和
			sum = 0
		} else {
			sum += current.Val
		}

		current = current.Next
	}

	return mergedHead
}
