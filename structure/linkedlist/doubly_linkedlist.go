package linkedlist

type doublynode struct {
	Val  int
	Prev *doublynode
	Next *doublynode
}

func NewDoublyNode(val int, prev, next *doublynode) *doublynode {
	return &doublynode{
		Val:  val,
		Prev: prev,
		Next: next,
	}
}

type DoublyLinkedList struct {
	size int
	Head *doublynode
}

func NewDoubly(items ...int) *DoublyLinkedList {
	list := DoublyLinkedList{
		size: 0,
		Head: nil,
	}
	curr := list.Head
	for _, val := range items {
		node := NewDoublyNode(val, nil, nil)
		if curr == nil {
			list.Head = node
			curr = list.Head
		} else {
			curr.Next = node
			node.Prev = curr
			curr = curr.Next
		}
		list.size++
	}

	return &list
}

func (l *DoublyLinkedList) Size() int {
	return l.size
}

func (l *DoublyLinkedList) InsertAt(index, val int) error {
	if index < 0 {
		return ErrInvalidIndex
	} else if index > l.size {
		return ErrIndexOutOfRange
	}
	node := NewDoublyNode(val, nil, nil)
	if index == 0 {
		if l.size == 0 {
			l.Head = node
		} else {
			node.Next = l.Head
			l.Head.Prev = node
			l.Head = node
		}
	} else {
		curr := l.Head
		for index > 1 {
			curr = curr.Next
			index--
		}
		node.Next = curr.Next
		if curr.Next != nil {
			curr.Next.Prev = node
		}
		node.Prev = curr
		curr.Next = node
	}
	l.size++

	return nil
}

func (l *DoublyLinkedList) BeginInsert(val int) {
	l.InsertAt(0, val)
}

func (l *DoublyLinkedList) LastInsert(val int) {
	l.InsertAt(l.size, val)
}

func (l *DoublyLinkedList) DeleteAt(index int) error {
	if index < 0 {
		return ErrInvalidIndex
	} else if index >= l.size {
		return ErrIndexOutOfRange
	}
	if index == 0 {
		if l.size == 1 {
			l.Head = nil
		} else {
			l.Head = l.Head.Next
			l.Head.Prev = nil
		}
	} else {
		curr := l.Head
		for index > 1 {
			curr = curr.Next
			index--
		}
		curr.Next = curr.Next.Next
		if curr.Next != nil {
			curr.Next.Prev = curr
		}
	}
	l.size--
	return nil
}

func (l *DoublyLinkedList) BeginDelete() {
	l.DeleteAt(0)
}

func (l *DoublyLinkedList) LastDelete() {
	l.DeleteAt(l.size - 1)
}

func (l *DoublyLinkedList) GetValueAt(index int) (int, error) {
	if index < 0 {
		return 0, ErrInvalidIndex
	} else if index >= l.size {
		return 0, ErrIndexOutOfRange
	}
	curr := l.Head
	for index > 0 {
		curr = curr.Next
		index--
	}

	return curr.Val, nil
}

func (l *DoublyLinkedList) Sort() {
	if l.size <= 1 {
		return
	}
	swapped := true
	for swapped {
		swapped = false
		for curr := l.Head; curr.Next != nil; curr = curr.Next {
			if curr.Val > curr.Next.Val {
				swapped = true
				curr.Val, curr.Next.Val = curr.Next.Val, curr.Val
			}
		}
	}
}

func (l *DoublyLinkedList) Reverse() {
	if l.size <= 1 {
		return
	}
	var prev, curr, next *doublynode
	curr = l.Head
	next = curr.Next
	for curr != nil {
		curr.Next = prev
		curr.Prev = next
		prev = curr
		curr = next
		if next != nil {
			next = next.Next
		}
	}
	l.Head = prev
}
