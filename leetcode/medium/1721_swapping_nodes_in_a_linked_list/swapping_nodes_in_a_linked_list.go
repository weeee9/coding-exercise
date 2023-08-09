package swappingnodesinalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	// 用來記錄鏈表的長度和當前位置的變數
	length := 0
	count := 1
	curr := head

	// 用來記錄從開頭數的第 k 個節點和從結尾數的第 k 個節點的指針
	var kthFromBeginning, kthFromEnd *ListNode

	// 遍歷鏈表以獲得從開頭數的第 k 個節點，並確定鏈表的長度
	for curr != nil {
		if count == k {
			kthFromBeginning = curr
		}
		curr = curr.Next
		count++
		length++
	}

	// 重設 count 並重新開始遍歷，以獲得從結尾數的第 k 個節點
	count = 1
	curr = head
	for curr != nil {
		if (count+k)-1 == length { // 使用先前確定的長度來識別從結尾數的第 k 個節點
			kthFromEnd = curr
			break
		}
		curr = curr.Next
		count++
	}

	// 交換從開頭數的第 k 個節點和從結尾數的第 k 個節點的值
	kthFromBeginning.Val, kthFromEnd.Val = kthFromEnd.Val, kthFromBeginning.Val

	return head
}
