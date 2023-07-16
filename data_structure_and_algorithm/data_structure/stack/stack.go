package stack

type Stack []int

func (s *Stack) Push(v int) {
	(*s) = append((*s), v)
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}

	index := len((*s)) - 1
	element := (*s)[index]
	(*s) = (*s)[:index]

	return element, true
}

func (s *Stack) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1
		return (*s)[index], true
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}
