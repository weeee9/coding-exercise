# Stack and Queue

## Stack

Stack 是一種重要的數據結構，它遵循 LIFO（Last In, First Out）的原則，也就是後進先出。可以把它想象成一疊盤子，新加入的盤子放在最上面，只有最上面的盤子能被移除。這種特性讓棧在處理某些問題上特別有用。

棧的主要操作包括：

1. Push：將元素放入棧的頂部。
2. Pop：移除棧頂的元素，並將其返回。
3. Peek 或 Top：返回棧頂的元素，但不將其從棧中移除。
4. IsEmpty：檢查棧是否為空。

```go
type Stack []int

// Push 添加元素到棧頂
func (s *Stack) Push(v int) {
	*s = append(*s, v)
}

// Pop 從棧頂移除元素並返回
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// Peek 返回棧頂的元素，但不將其從棧中移除
func (s Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(s) - 1
		return s[index], true
	}
}

// IsEmpty 檢查棧是否為空
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
```

## Queue

隊列（Queue）是一種數據結構，它遵循 FIFO（First-In-First-Out，先進先出）的原則。你可以將它想象成一種管道，新的元素從一端加入，而已存在的元素則從另一端取出。

隊列的主要操作有：

1. Enqueue：將元素添加到隊列的尾部。
2. Dequeue：將元素從隊列的頭部移除並返回。
3. Front：返回隊列的第一個元素，但不移除它。
4. IsEmpty：檢查隊列是否為空。

```go
type Queue []int

// Enqueue 添加元素到隊列尾部
func (q *Queue) Enqueue(v int) {
	*q = append(*q, v)
}

// Dequeue 從隊列頭部移除元素並返回
func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}

// Front 返回隊列的第一個元素，但不移除它
func (q *Queue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return (*q)[0], true
}

// IsEmpty 檢查隊列是否為空
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
```
