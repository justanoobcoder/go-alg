package set

type Set interface {
	Size() int
	GetItems() []int
	Contains(n int) bool
	Add(n int)
	Delete(n int)
	IsSubsetOf(other Set) bool
	Union(other Set) Set
	Intersection(other Set) Set
	Difference(other Set) Set
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

func (s *set) Union(other Set) Set {
	union := New()
	for _, n := range s.GetItems() {
		union.Add(n)
	}
	for _, n := range other.GetItems() {
		union.Add(n)
	}
	return union
}

func (s *set) Intersection(other Set) Set {
	result := New()
	var minSet, maxSet Set
	if s.Size() < other.Size() {
		minSet = s
		maxSet = other
	} else {
		minSet = other
		maxSet = s
	}
	for _, n := range minSet.GetItems() {
		if maxSet.Contains(n) {
			result.Add(n)
		}
	}

	return result
}

func (s *set) Difference(other Set) Set {
	result := New()
	for n := range s.elements {
		if !other.Contains(n) {
			result.Add(n)
		}
	}

	return result
}
