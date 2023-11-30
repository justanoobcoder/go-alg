package set

type Set interface {
	Size() int
	GetItems() []int
	Contains(n int) bool
	Add(n int)
	Delete(n int)
	IsSubsetOf(other Set) bool
}

type set struct {
	elements map[int]bool
}

func New(items ...int) Set {
	s := set{
		elements: make(map[int]bool),
	}
	for _, n := range items {
		s.Add(n)
	}
	return &s
}

func (s *set) GetItems() []int {
	items := make([]int, 0)
	for k := range s.elements {
		items = append(items, k)
	}

	return items
}

func (s *set) Size() int {
	return len(s.elements)
}

func (s *set) Contains(n int) bool {
	_, ok := s.elements[n]
	return ok
}

func (s *set) Add(n int) {
	s.elements[n] = true
}

func (s *set) Delete(n int) {
	delete(s.elements, n)
}

func (s *set) IsSubsetOf(other Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	items := s.GetItems()
	for i := 0; i < len(items); i++ {
		if !other.Contains(items[i]) {
			return false
		}
	}

	return true
}
