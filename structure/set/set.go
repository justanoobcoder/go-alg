package set

type Set interface {
	Size() int
	Contains(n int) bool
	Add(n int)
	Delete(n int)
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
