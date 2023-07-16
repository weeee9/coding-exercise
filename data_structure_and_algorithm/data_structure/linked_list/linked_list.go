package linkedlist

// 節點
type Node struct {
	data int
	next *Node
}

// 鏈表
type LinkedList struct {
	head *Node
	tail *Node
	size int
}

// 在position插入data
func (l *LinkedList) Insert(position, data int) {
	// position超出範圍，則不插入
	if position > l.size {
		return
	}

	newNode := &Node{data, nil}

	switch {
	// 從頭插入
	case position == 0:
		newNode.next = l.head
		l.head = newNode
		if l.tail == nil {
			l.tail = newNode
		}
	// 從尾插入
	case position == l.size:
		l.tail.next = newNode
		l.tail = newNode
	// 從中插入
	default:
		prev := l.head
		for i := 0; i < position; i++ {
			prev = prev.next
		}

		next := prev.next
		newNode.next = next
		prev.next = newNode
	}
	l.size++
}

// 搜尋data的位置
func (l *LinkedList) Search(data int) int {
	current := l.head
	for i := 0; i < l.size; i++ {
		if current.data == data {
			return i
		}
		current = current.next
	}

	return -1
}

// 更新position的data
func (l *LinkedList) Update(position, data int) {
	if position > l.size {
		return
	}

	current := l.head
	for i := 0; i < position; i++ {
		current = current.next
	}

	current.data = data
}

// 刪除data
func (l *LinkedList) Delete(data int) {
	if l.size == 0 {
		return
	}

	// 刪除的是頭節點
	if l.head.data == data {
		l.head = l.head.next
		if l.head == nil { // 若刪除後鏈表為空，則尾節點也要設為nil
			l.tail = nil
		}
		return
	}

	// 刪除的是中間或尾節點
	for prev, node := l.head, l.head.next; node != nil; prev, node = node, node.next {
		if node.data == data {
			prev.next = node.next
			if node == l.tail { // 若刪除的是尾節點，則更新尾節點
				l.tail = prev
			}
			return
		}
	}

	l.size--
}
