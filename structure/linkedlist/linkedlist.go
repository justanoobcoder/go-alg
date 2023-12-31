package linkedlist

import "errors"

type Linkedlist interface {
	Size() int
	GetValueAt(index int) (int, error)
	BeginInsert(n int)
	BeginDelete()
	LastInsert(n int)
	LastDelete()
	InsertAt(index, value int) error
	DeleteAt(index int) error
	Sort()
	Reverse()
}

var (
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrInvalidIndex    = errors.New("invalid index")
)
