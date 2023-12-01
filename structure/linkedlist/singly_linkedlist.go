package linkedlist

import "errors"

type Node struct {
	Val  int
	Next *Node
}

func newNode(n int, next *Node) *Node {
	return &Node{
		Val:  n,
		Next: next,
	}
}

type SinglyLinkedlist struct {
	size int
	Head *Node
}

func NewSingly(items ...int) *SinglyLinkedlist {
	list := SinglyLinkedlist{
		size: 0,
		Head: nil,
	}
	curr := list.Head
	for _, n := range items {
		if curr == nil {
			list.Head = newNode(n, nil)
			curr = list.Head
		} else {
			curr.Next = newNode(n, nil)
			curr = curr.Next
		}
		list.size++
	}
	return &list
}

func (l *SinglyLinkedlist) Size() int {
	return l.size
}

func (l *SinglyLinkedlist) BeginInsert(n int) {
	if l.size == 0 {
		l.Head = newNode(n, nil)
	} else {
		node := newNode(n, l.Head)
		l.Head = node
	}
	l.size++
}

func (l *SinglyLinkedlist) LastInsert(n int) {
	if l.size == 0 {
		l.Head = newNode(n, nil)
	} else {
		node := newNode(n, nil)
		curr := l.Head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}
	l.size++
}

func (l *SinglyLinkedlist) BeginDelete() {
	if l.size == 0 {
		return
	} else if l.size == 1 {
		l.Head = nil
	} else {
		dummy := l.Head
		l.Head = l.Head.Next
		dummy.Next = nil
	}
	l.size--
}

func (l *SinglyLinkedlist) LastDelete() {
	if l.size == 0 {
		return
	} else if l.size == 1 {
		l.Head = nil
	} else {
		curr := l.Head
		for curr.Next.Next != nil {
			curr = curr.Next
		}
		curr.Next = nil
	}
	l.size--
}

func (l *SinglyLinkedlist) Sort() {
	if l.size <= 1 {
		return
	}
	swapped := true
	for swapped {
		swapped = false
		for curr := l.Head; curr.Next.Next != nil; curr = curr.Next {
			if curr.Val > curr.Next.Val {
				curr.Val, curr.Next.Val = curr.Next.Val, curr.Val
				swapped = true
			}
		}
	}
}

func (l *SinglyLinkedlist) Reverse() {
	if l.size <= 1 {
		return
	}
	reverseList := NewSingly()
	for curr := l.Head; curr != nil; curr = curr.Next {
		reverseList.BeginInsert(curr.Val)
	}
	l.Head = reverseList.Head
}

func (l *SinglyLinkedlist) InsertAt(index, value int) error {
	if index < 0 {
		return errors.New("index is invalid")
	} else if index > l.size {
		return errors.New("index is out of range")
	}

	if index == 0 {
		l.BeginInsert(value)
	} else if index == l.size {
		l.LastInsert(value)
	} else {
		curr := l.Head
		for index > 1 {
			curr = curr.Next
			index--
		}
		node := newNode(value, curr.Next)
		curr.Next = node
		l.size++
	}

	return nil
}

func (l *SinglyLinkedlist) DeleteAt(index int) error {
	if index < 0 {
		return errors.New("index is invalid")
	} else if index >= l.size {
		return errors.New("index is out of range")
	}

	if index == 0 {
		l.BeginDelete()
	} else if index == l.size-1 {
		l.LastDelete()
	} else {
		curr := l.Head
		for index > 1 {
			curr = curr.Next
			index--
		}
		curr.Next = curr.Next.Next
		l.size--
	}

	return nil
}
