package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TESTS

func TestNew(t *testing.T) {
	items := []int{1, 2, 3, 4}
	s := New(items...)
	assert.Equal(t, len(items), s.Size())
	for _, n := range items {
		assert.True(t, s.Contains(n))
	}
	assert.False(t, s.Contains(3829))
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
			s := New(data...)
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
			s := New(data...)
			s.Delete(test.input)
			assert.Equal(t, len(test.expected), s.Size())
			assert.False(t, s.Contains(test.input))
		})
	}
}
