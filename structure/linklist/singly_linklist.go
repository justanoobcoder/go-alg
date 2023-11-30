package linklist

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

type SinglyLinklist struct {
	size int
	Head *Node
}

func NewSingly(items ...int) *SinglyLinklist {
	list := SinglyLinklist{
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

func (l *SinglyLinklist) Size() int {
	return l.size
}

func (l *SinglyLinklist) BeginInsert(n int) {
	if l.size == 0 {
		l.Head = newNode(n, nil)
	} else {
		node := newNode(n, l.Head)
		l.Head = node
	}
	l.size++
}

func (l *SinglyLinklist) LastInsert(n int) {
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
