package linkedlist_test

import (
	"reflect"
	"testing"

	"github.com/justanoobcoder/go-alg/structure/linkedlist"
)

// TESTS

func TestNewDoublyLinkedList(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "create an empty doubly linkedlist",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "create a doubly linkedlist with some items",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			got := make([]int, 0)
			list := linkedlist.NewDoubly(test.input...)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			if len(test.expected) != list.Size() {
				t.Errorf("want size %v, got size %v", len(test.expected), list.Size())
			}
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("want %v, got %v", test.expected, got)
			}
		})
	}
}
