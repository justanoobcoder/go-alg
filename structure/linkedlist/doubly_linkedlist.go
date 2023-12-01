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
