package removenthnodefromendoflist

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	// 計算鏈結串列的總節點數
	totalLen := 0
	current := head
	for current != nil {
		totalLen++
		current = current.Next
	}

	// 如果要刪除的節點是第一個節點，直接將 head 指向第二個節點，並返回 head
	if n == totalLen {
		head = head.Next
		return head
	}

	// 找到倒數第 n 個節點，並將其從鏈結串列中移除
	current, previous := head, head
	idx := 0
	for current != nil {
		idx++

		// n 從 totalLen+1 開始計算，找到倒數第 n 個節點，將其從鏈結串列中移除
		if idx == totalLen+1-n {
			previous.Next = current.Next
		}

		previous = current
		current = current.Next
	}

	return head
}
