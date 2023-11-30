package linklist_test

import (
	"reflect"
	"testing"

	"github.com/justanoobcoder/go-alg/structure/linklist"
	"github.com/stretchr/testify/assert"
)

// TESTS

func TestNew(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "create an empty singly linklist",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "create a singly linklist with some items",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			got := make([]int, 0)
			list := linklist.NewSingly(test.input...)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			assert.Equal(t, len(test.expected), list.Size())
			assert.ElementsMatchf(t, test.expected, got, "want %v, got %v",
				test.expected, got)
		})
	}
}

func TestBeginInsert(t *testing.T) {
	data := []int{1, 2, 3}
	want := []int{3, 2, 1}
	list := linklist.NewSingly()
	for _, n := range data {
		list.BeginInsert(n)
	}
	got := make([]int, 0)
	for curr := list.Head; curr != nil; curr = curr.Next {
		got = append(got, curr.Val)
	}
	assert.Equal(t, len(want), list.Size(),
		"want %v, got %v", len(want), list.Size())
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

func TestLastInsert(t *testing.T) {
	want := []int{1, 2, 3}
	list := linklist.NewSingly()
	for _, n := range want {
		list.LastInsert(n)
	}
	got := make([]int, 0)
	for curr := list.Head; curr != nil; curr = curr.Next {
		got = append(got, curr.Val)
	}
	assert.Equal(t, len(want), list.Size(),
		"want %v, got %v", len(want), list.Size())
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %v, got %v", want, got)
	}
}

// BENCHMARKS

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linklist.NewSingly(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkBeginInsert(b *testing.B) {
	list := linklist.NewSingly()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.BeginInsert(1)
	}
}

func BenchmarkLastInsert(b *testing.B) {
	list := linklist.NewSingly()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.LastInsert(1)
	}
}
