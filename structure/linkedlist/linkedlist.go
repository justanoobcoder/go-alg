package linkedlist

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
