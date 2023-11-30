package set_test

import (
	"testing"

	"github.com/justanoobcoder/go-alg/structure/set"
	"github.com/stretchr/testify/assert"
)

// TESTS

func TestNew(t *testing.T) {
	items := []int{1, 2, 3, 4}
	s := set.New(items...)
	assert.Equal(t, len(items), s.Size())
	for _, n := range items {
		assert.True(t, s.Contains(n))
	}
	assert.False(t, s.Contains(3829))
}

func TestGetItems(t *testing.T) {
	data := []int{1, 2, 3, 4}
	s := set.New(data...)
	items := s.GetItems()
	assert.Equal(t, len(data), len(items))
	for _, n := range data {
		assert.Contains(t, items, n)
	}
}

func TestContains(t *testing.T) {
	data := []int{1, 2, 3, 4}
	table := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "contains existing item",
			input:    3,
			expected: true,
		},
		{
			name:     "contains not existing item",
			input:    5,
			expected: false,
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			s := set.New(data...)
			assert.Equal(t, test.expected, s.Contains(test.input))
		})
	}
}

func TestSize(t *testing.T) {
	data := []int{1, 2, 3, 4}
	s := set.New(data...)
	assert.Equal(t, len(data), s.Size())
}

func TestAdd(t *testing.T) {
	data := []int{1, 2, 3, 4}
	table := []struct {
		name     string
		input    int
		expected []int
	}{
		{
			name:     "add new item",
			input:    5,
			expected: append(data, 5),
		},
		{
			name:     "add existing item",
			input:    3,
			expected: data,
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			s := set.New(data...)
			s.Add(test.input)
			assert.Equal(t, len(test.expected), s.Size())
			for _, n := range test.expected {
				assert.True(t, s.Contains(n))
			}
		})
	}
}

func TestDelete(t *testing.T) {
	data := []int{1, 2, 3, 4}
	table := []struct {
		name     string
		input    int
		expected []int
	}{
		{
			name:     "delete existing item",
			input:    3,
			expected: []int{1, 2, 4},
		},
		{
			name:     "delete not existing item",
			input:    5,
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			s := set.New(data...)
			s.Delete(test.input)
			assert.Equal(t, len(test.expected), s.Size())
			assert.False(t, s.Contains(test.input))
		})
	}
}

func TestIsSubsetOf(t *testing.T) {
	s1, s2, s3 := set.New(1, 2, 3), set.New(1, 2, 3, 4, 5), set.New(9, 10, 5, 8)
	assert.False(t, s2.IsSubsetOf(s1))
	assert.True(t, s1.IsSubsetOf(s2))
	assert.False(t, s1.IsSubsetOf(s3))
}

func TestUnion(t *testing.T) {
	table := []struct {
		name     string
		s1       set.Set
		s2       set.Set
		expected set.Set
	}{
		{
			name:     "union of disjoint sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(4, 5, 6),
			expected: set.New(1, 2, 3, 4, 5, 6),
		},
		{
			name:     "union of overlapping sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(3, 4, 5),
			expected: set.New(1, 2, 3, 4, 5),
		},
		{
			name:     "union of same sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(1, 2, 3),
			expected: set.New(1, 2, 3),
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			s := test.s1.Union(test.s2)
			assert.Equal(t, test.expected.Size(), s.Size())
			for _, n := range test.expected.GetItems() {
				assert.True(t, s.Contains(n))
			}
		})
	}
}

func TestIntersection(t *testing.T) {
	table := []struct {
		name     string
		s1       set.Set
		s2       set.Set
		expected set.Set
	}{
		{
			name:     "intersection of disjoint sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(4, 5, 6),
			expected: set.New(),
		},
		{
			name:     "intersection of overlapping sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(3, 4, 5),
			expected: set.New(3),
		},
		{
			name:     "intersection of same sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(1, 2, 3),
			expected: set.New(1, 2, 3),
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			s := test.s1.Intersection(test.s2)
			assert.Equal(t, test.expected.Size(), s.Size())
			for _, n := range test.expected.GetItems() {
				assert.True(t, s.Contains(n))
			}
		})
	}
}

func TestDifference(t *testing.T) {
	table := []struct {
		name     string
		s1       set.Set
		s2       set.Set
		expected set.Set
	}{
		{
			name:     "difference of disjoint sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(4, 5, 6),
			expected: set.New(1, 2, 3),
		},
		{
			name:     "difference of overlapping sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(3, 4, 5),
			expected: set.New(1, 2),
		},
		{
			name:     "difference of same sets",
			s1:       set.New(1, 2, 3),
			s2:       set.New(1, 2, 3),
			expected: set.New(),
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			s := test.s1.Difference(test.s2)
			assert.Equal(t, test.expected.Size(), s.Size())
			for _, n := range test.expected.GetItems() {
				assert.True(t, s.Contains(n))
			}
		})
	}
}
