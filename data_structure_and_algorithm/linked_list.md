# Linked List

![Image.jpeg](https://i0.wp.com/www.tutorialspoint.com/data_structures_algorithms/images/linked_list.jpg?w=1170&ssl=1)

一個鏈表是一種序列數據結構，其中每個元素都包含兩部分的信息：一部分是元素本身的數據，另一部分則是指向下一個元素的參考（通常稱為 "pointer" 或者 "link"）。每個元素（或稱為節點）通常由一個數據欄位和一個指針欄位組成。

這種數據結構的主要優點是它允許在不重新排列整個數據結構的情況下插入和刪除元素。同時，相比於陣列等其它數據結構，鏈表在內存空間的使用上更具有彈性，因為它不需要在創建時確定其大小，且可以根據需要動態地增加或減少節點。

但鏈表也有其缺點。例如，對於鏈表，我們無法直接訪問某個元素（如，透過索引訪問，就像在陣列中那樣）。要訪問鏈表中的某個元素，必須從第一個元素（通常稱為 "頭節點"）開始，並沿著鏈表逐一訪問每個元素，直到找到所需的元素。這就導致訪問元素的時間複雜度為 O(n)。

鏈表可以有多種變體，包括：

- 單向鏈表（每個節點只有一個指向下一個節點的指針）

![Image.png](https://upload.wikimedia.org/wikipedia/commons/thumb/6/6d/Singly-linked-list.svg/408px-Singly-linked-list.svg.png)

- 雙向鏈表（每個節點有兩個指針，一個指向前一個節點，一個指向下一個節點）

![Image.png](https://upload.wikimedia.org/wikipedia/commons/thumb/5/5e/Doubly-linked-list.svg/610px-Doubly-linked-list.svg.png)

- 循環鏈表（最後一個節點指向第一個節點，形成一個循環）等。這些變體有各自的用途和優點。

![Image.png](https://upload.wikimedia.org/wikipedia/commons/thumb/d/df/Circularly-linked-list.svg/350px-Circularly-linked-list.svg.png)

## Linked List 的操作

```go
type Node struct {
  data int
  next *Node
}

type LinkedList struct {
  head *Node
  tail *Node
  size int
}

func (l *LinkedList) Insert((position, data int) {
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

// 在鏈表中搜索特定的數據
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

// 在鏈表中刪除特定的數據
func (l *LinkedList) Delete(data int) {
	if l.head == nil { // 若鏈表為空，則直接返回
		return
	}

	// 若首節點就是要刪除的節點
	if l.head.data == data {
		l.head = l.head.next
		if l.head == nil { // 若刪除後鏈表為空，則尾節點也要設為nil
			l.tail = nil
		}
		return
	}

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
```