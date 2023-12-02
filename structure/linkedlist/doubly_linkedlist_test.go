package linkedlist_test

import (
	"errors"
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

func TestDoublyInsertAt(t *testing.T) {
	table := []struct {
		name        string
		input       []int
		index       int
		val         int
		expected    []int
		expectedErr error
	}{
		{
			name:        "insert at the beginning of an empty doubly linkedlist",
			input:       []int{},
			index:       0,
			val:         1,
			expected:    []int{1},
			expectedErr: nil,
		},
		{
			name:        "insert at the beginning of a non-empty doubly linkedlist",
			input:       []int{2, 3, 4, 5},
			index:       0,
			val:         1,
			expected:    []int{1, 2, 3, 4, 5},
			expectedErr: nil,
		},
		{
			name:        "insert at the end of a non-empty doubly linkedlist",
			input:       []int{1, 2, 3, 4},
			index:       4,
			val:         5,
			expected:    []int{1, 2, 3, 4, 5},
			expectedErr: nil,
		},
		{
			name:        "insert at the middle of a non-empty doubly linkedlist",
			input:       []int{1, 2, 4, 5},
			index:       2,
			val:         3,
			expected:    []int{1, 2, 3, 4, 5},
			expectedErr: nil,
		},
		{
			name:        "insert at the end of an empty doubly linkedlist",
			input:       []int{},
			index:       1,
			val:         1,
			expected:    []int{},
			expectedErr: linkedlist.ErrIndexOutOfRange,
		},
		{
			name:        "insert with negative index",
			input:       []int{},
			index:       -1,
			val:         1,
			expected:    []int{},
			expectedErr: linkedlist.ErrInvalidIndex,
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			err := list.InsertAt(test.index, test.val)
			got := make([]int, 0)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			if test.expectedErr == nil && err != nil {
				t.Errorf("want no error, got %v", err)
			}
			if test.expectedErr != nil && err == nil {
				t.Errorf("want error %v, got no error", test.expectedErr)
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

func TestDoublyBeginInsert(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		val      int
		expected []int
	}{
		{
			name:     "insert at the beginning of an empty doubly linkedlist",
			input:    []int{},
			val:      1,
			expected: []int{1},
		},
		{
			name:     "insert at the beginning of a non-empty doubly linkedlist",
			input:    []int{2, 3, 4, 5},
			val:      1,
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			list.BeginInsert(test.val)
			got := make([]int, 0)
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

func TestDoublyLastInsert(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		val      int
		expected []int
	}{
		{
			name:     "insert at the end of an empty doubly linkedlist",
			input:    []int{},
			val:      1,
			expected: []int{1},
		},
		{
			name:     "insert at the end of a non-empty doubly linkedlist",
			input:    []int{1, 2, 3, 4},
			val:      5,
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			list.LastInsert(test.val)
			got := make([]int, 0)
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

func TestDoublyDeleteAt(t *testing.T) {
	table := []struct {
		name        string
		input       []int
		index       int
		expected    []int
		expectedErr error
	}{
		{
			name:        "delete at the beginning of a non-empty doubly linkedlist",
			input:       []int{1, 2, 3, 4},
			index:       0,
			expected:    []int{2, 3, 4},
			expectedErr: nil,
		},
		{
			name:        "delete at the end of a non-empty doubly linkedlist",
			input:       []int{1, 2, 3, 4},
			index:       3,
			expected:    []int{1, 2, 3},
			expectedErr: nil,
		},
		{
			name:        "delete at the middle of a non-empty doubly linkedlist",
			input:       []int{1, 2, 3, 4},
			index:       2,
			expected:    []int{1, 2, 4},
			expectedErr: nil,
		},
		{
			name:        "delete at the beginning of an empty doubly linkedlist",
			input:       []int{},
			index:       0,
			expected:    []int{},
			expectedErr: linkedlist.ErrIndexOutOfRange,
		},
		{
			name:        "delete with negative index",
			input:       []int{},
			index:       -1,
			expected:    []int{},
			expectedErr: linkedlist.ErrInvalidIndex,
		},
		{
			name:        "delete with index out of range",
			input:       []int{1, 2, 3, 4},
			index:       4,
			expected:    []int{1, 2, 3, 4},
			expectedErr: linkedlist.ErrIndexOutOfRange,
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			err := list.DeleteAt(test.index)
			got := make([]int, 0)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			if test.expectedErr == nil && err != nil {
				t.Errorf("want no error, got %v", err)
			}
			if test.expectedErr != nil && err == nil {
				t.Errorf("want error %v, got no error", test.expectedErr)
			}
			if test.expectedErr != nil && !errors.Is(err, test.expectedErr) {
				t.Errorf("want error %v, got error %v", test.expectedErr, err)
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

func TestDoublyBeginDelete(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "delete at the beginning of a non-empty doubly linkedlist",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 3, 4},
		},
		{
			name:     "delete at the beginning of an empty doubly linkedlist",
			input:    []int{},
			expected: []int{},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			list.BeginDelete()
			got := make([]int, 0)
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

func TestDoublyLastDelete(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "delete at the end of a non-empty doubly linkedlist",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3},
		},
		{
			name:     "delete at the end of an empty doubly linkedlist",
			input:    []int{},
			expected: []int{},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			list.LastDelete()
			got := make([]int, 0)
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

func TestDoublyGetValueAt(t *testing.T) {
	table := []struct {
		name        string
		input       []int
		index       int
		expected    int
		expectedErr error
	}{
		{
			name:        "get value at the beginning of a non-empty doubly linkedlist",
			input:       []int{1, 2, 3, 4},
			index:       0,
			expected:    1,
			expectedErr: nil,
		},
		{
			name:        "get value at the end of a non-empty doubly linkedlist",
			input:       []int{1, 2, 3, 4},
			index:       3,
			expected:    4,
			expectedErr: nil,
		},
		{
			name:        "get value at the middle of a non-empty doubly linkedlist",
			input:       []int{1, 2, 3, 4},
			index:       2,
			expected:    3,
			expectedErr: nil,
		},
		{
			name:        "get value at the beginning of an empty doubly linkedlist",
			input:       []int{},
			index:       0,
			expected:    0,
			expectedErr: linkedlist.ErrIndexOutOfRange,
		},
		{
			name:        "get value with negative index",
			input:       []int{},
			index:       -1,
			expected:    0,
			expectedErr: linkedlist.ErrInvalidIndex,
		},
		{
			name:        "get value with index out of range",
			input:       []int{1, 2, 3, 4},
			index:       4,
			expected:    0,
			expectedErr: linkedlist.ErrIndexOutOfRange,
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			got, err := list.GetValueAt(test.index)
			if test.expectedErr == nil && err != nil {
				t.Errorf("want no error, got %v", err)
			}
			if test.expectedErr != nil && err == nil {
				t.Errorf("want error %v, got no error", test.expectedErr)
			}
			if test.expectedErr != nil && !errors.Is(err, test.expectedErr) {
				t.Errorf("want error %v, got error %v", test.expectedErr, err)
			}
			if test.expected != got {
				t.Errorf("want %v, got %v", test.expected, got)
			}
		})
	}
}

func TestDoublySort(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "sort a non-empty doubly linkedlist",
			input:    []int{4, 2, 3, 1},
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "sort an empty doubly linkedlist",
			input:    []int{},
			expected: []int{},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			list.Sort()
			got := make([]int, 0)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("want %v, got %v", test.expected, got)
			}
		})
	}
}

func TestDoublyReverse(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "reverse a non-empty doubly linkedlist",
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "reverse an empty doubly linkedlist",
			input:    []int{},
			expected: []int{},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewDoubly(test.input...)
			list.Reverse()
			reversedList := make([]int, 0)
			originList := make([]int, 0)
			curr := list.Head
			for curr != nil && curr.Next != nil {
				reversedList = append(reversedList, curr.Val)
				curr = curr.Next
			}
			if curr != nil {
				reversedList = append(reversedList, curr.Val)
			}
			for ; curr != nil; curr = curr.Prev {
				originList = append(originList, curr.Val)
			}
			if !reflect.DeepEqual(test.input, originList) {
				t.Errorf("want %v, originList %v", test.input, originList)
			}
			if !reflect.DeepEqual(test.expected, reversedList) {
				t.Errorf("want %v, reversedList %v", test.expected, reversedList)
			}
		})
	}
}
