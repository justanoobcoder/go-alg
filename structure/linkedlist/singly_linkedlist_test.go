package linkedlist_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/justanoobcoder/go-alg/structure/linkedlist"
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
			name:     "create an empty singly linkedlist",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "create a singly linkedlist with some items",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			got := make([]int, 0)
			list := linkedlist.NewSingly(test.input...)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			assert.Equal(t, len(test.expected), list.Size())
			assert.ElementsMatchf(t, test.expected, got, "want %v, got %v",
				test.expected, got)
		})
	}
}

func TestGetValueAt(t *testing.T) {
	table := []struct {
		name      string
		input     []int
		index     int
		expected  int
		expectErr error
	}{
		{
			name:      "get value at index 0 of an empty linkedlist",
			input:     []int{},
			index:     0,
			expected:  0,
			expectErr: errors.New("index is out of range"),
		},
		{
			name:      "get value at index 0 of a linkedlist that has 1 node",
			input:     []int{1},
			index:     0,
			expected:  1,
			expectErr: nil,
		},
		{
			name:      "get value at index 0 of a linkedlist that has 2 nodes",
			input:     []int{1, 2},
			index:     0,
			expected:  1,
			expectErr: nil,
		},
		{
			name:      "get value at index 1 of a linkedlist that has 2 nodes",
			input:     []int{1, 2},
			index:     1,
			expected:  2,
			expectErr: nil,
		},
		{
			name:      "get value at index 2 of a linkedlist that has 2 nodes",
			input:     []int{1, 2},
			index:     2,
			expected:  0,
			expectErr: errors.New("index is out of range"),
		},
		{
			name:      "get value at index -1 of a linkedlist that has 2 nodes",
			input:     []int{1, 2},
			index:     -1,
			expected:  0,
			expectErr: errors.New("index is invalid"),
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewSingly(test.input...)
			got, err := list.GetValueAt(test.index)
			if test.expectErr == nil && err != nil {
				t.Errorf("expected no err, actual %q", err)
			}
			if test.expectErr != nil && err == nil {
				t.Errorf("expected err %q, actual no err", test.expectErr)
			}
			if test.expectErr != nil && err.Error() != test.expectErr.Error() {
				t.Errorf("expected error message %q, acutal error message %q", test.expectErr, err)
			}
			if test.expected != got {
				t.Errorf("want %v, got %v", test.expected, got)
			}
		})
	}
}

func TestBeginInsert(t *testing.T) {
	data := []int{1, 2, 3}
	want := []int{3, 2, 1}
	list := linkedlist.NewSingly()
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
	list := linkedlist.NewSingly()
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

func TestBeginDelete(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "delete begin node of an empty linkedlist",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "delete begin node of a linkedlist that has one node",
			input:    []int{3},
			expected: []int{},
		},
		{
			name:     "delete begin node of a linkedlist that has some nodes",
			input:    []int{1, 2, 3, 4},
			expected: []int{2, 3, 4},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewSingly(test.input...)
			list.BeginDelete()
			got := make([]int, 0)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			if wantSize, gotSize := len(test.expected), list.Size(); wantSize != gotSize {
				t.Errorf("want list size %v, got list size %v", wantSize, gotSize)
			}
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("want %v, got %v", test.expected, got)
			}
		})
	}
}

func TestLastDelete(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "delete last node of an empty linkedlist",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "delete last node of a linkedlist that has one node",
			input:    []int{3},
			expected: []int{},
		},
		{
			name:     "delete last node of a linkedlist that has some nodes",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewSingly(test.input...)
			list.LastDelete()
			got := make([]int, 0)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			if wantSize, gotSize := len(test.expected), list.Size(); wantSize != gotSize {
				t.Errorf("want list size %v, got list size %v", wantSize, gotSize)
			}
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("want %v, got %v", test.expected, got)
			}
		})
	}
}

func TestSort(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "sort an empty linkedlist",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "sort a linkedlist that has one node",
			input:    []int{3},
			expected: []int{3},
		},
		{
			name:     "sort a linkedlist that has some nodes",
			input:    []int{1, 3, 2, 4},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewSingly(test.input...)
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

func TestReverse(t *testing.T) {
	table := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "reverse an empty linkedlist",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "reverse a linkedlist that has one node",
			input:    []int{3},
			expected: []int{3},
		},
		{
			name:     "reverse a linkedlist that has some nodes",
			input:    []int{1, 3, 2, 4},
			expected: []int{4, 2, 3, 1},
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewSingly(test.input...)
			list.Reverse()
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

func TestInsertAt(t *testing.T) {
	table := []struct {
		name      string
		input     []int
		index     int
		value     int
		expected  []int
		expectErr error
	}{
		{
			name:      "insert at index 0 of an empty linkedlist",
			input:     []int{},
			index:     0,
			value:     1,
			expected:  []int{1},
			expectErr: nil,
		},
		{
			name:      "insert at index 0 of a linkedlist that has 3 nodes",
			input:     []int{1, 2, 3},
			index:     0,
			value:     4,
			expected:  []int{4, 1, 2, 3},
			expectErr: nil,
		},
		{
			name:      "insert at index 1 of a linkedlist that has 3 nodes",
			input:     []int{1, 2, 3},
			index:     1,
			value:     4,
			expected:  []int{1, 4, 2, 3},
			expectErr: nil,
		},
		{
			name:      "insert at index 2 of a linkedlist that has 3 nodes",
			input:     []int{1, 2, 3},
			index:     2,
			value:     4,
			expected:  []int{1, 2, 4, 3},
			expectErr: nil,
		},
		{
			name:      "insert at index 3 of a linkedlist that has 3 nodes",
			input:     []int{1, 2, 3},
			index:     3,
			value:     4,
			expected:  []int{1, 2, 3, 4},
			expectErr: nil,
		},
		{
			name:      "insert at index 4 of a linkedlist that has 3 nodes",
			input:     []int{1, 2, 3},
			index:     4,
			value:     4,
			expected:  []int{1, 2, 3},
			expectErr: errors.New("index is out of range"),
		},
		{
			name:      "insert at index -1 of a linkedlist that has 3 nodes",
			input:     []int{1, 2, 3},
			index:     -1,
			value:     4,
			expected:  []int{1, 2, 3},
			expectErr: errors.New("index is invalid"),
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewSingly(test.input...)
			err := list.InsertAt(test.index, test.value)
			got := make([]int, 0)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			if test.expectErr == nil && err != nil {
				t.Errorf("expected no err, actual %q", err)
			}
			if test.expectErr != nil && err == nil {
				t.Errorf("expected err %q, actual no err", test.expectErr)
			}
			if test.expectErr != nil && err.Error() != test.expectErr.Error() {
				t.Errorf("expected error message %q, acutal error message %q", test.expectErr, err)
			}
			if wantSize, gotSize := len(test.expected), list.Size(); wantSize != gotSize {
				t.Errorf("want size %v, got size %v", wantSize, gotSize)
			}
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("want %v, got %v", test.expected, got)
			}
		})
	}
}

func TestDeleteAt(t *testing.T) {
	table := []struct {
		name      string
		input     []int
		index     int
		expected  []int
		expectErr error
	}{
		{
			name:      "delete at index 0 of an empty linkedlist",
			input:     []int{},
			index:     0,
			expected:  []int{},
			expectErr: errors.New("index is out of range"),
		},
		{
			name:      "delete at index 0 of a linkedlist that has 1 node",
			input:     []int{1},
			index:     0,
			expected:  []int{},
			expectErr: nil,
		},
		{
			name:      "delete at index 0 of a linkedlist that has 2 nodes",
			input:     []int{1, 2},
			index:     0,
			expected:  []int{2},
			expectErr: nil,
		},
		{
			name:      "delete at index 1 of a linkedlist that has 2 nodes",
			input:     []int{1, 2},
			index:     1,
			expected:  []int{1},
			expectErr: nil,
		},
		{
			name:      "delete at index 2 of a linkedlist that has 5 nodes",
			input:     []int{1, 2, 3, 4, 5},
			index:     2,
			expected:  []int{1, 2, 4, 5},
			expectErr: nil,
		},
		{
			name:      "delete at index 2 of a linkedlist that has 2 nodes",
			input:     []int{1, 2},
			index:     2,
			expected:  []int{1, 2},
			expectErr: errors.New("index is out of range"),
		},
		{
			name:      "delete at index -1 of a linkedlist that has 2 nodes",
			input:     []int{1, 2},
			index:     -1,
			expected:  []int{1, 2},
			expectErr: errors.New("index is invalid"),
		},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			list := linkedlist.NewSingly(test.input...)
			err := list.DeleteAt(test.index)
			got := make([]int, 0)
			for curr := list.Head; curr != nil; curr = curr.Next {
				got = append(got, curr.Val)
			}
			if test.expectErr == nil && err != nil {
				t.Errorf("expected no err, actual %q", err)
			}
			if test.expectErr != nil && err == nil {
				t.Errorf("expected err %q, actual no err", test.expectErr)
			}
			if test.expectErr != nil && err.Error() != test.expectErr.Error() {
				t.Errorf("expected error message %q, acutal error message %q", test.expectErr, err)
			}
			if wantSize, gotSize := len(test.expected), list.Size(); wantSize != gotSize {
				t.Errorf("want size %v, got size %v", wantSize, gotSize)
			}
			if !reflect.DeepEqual(test.expected, got) {
				t.Errorf("want %v, got %v", test.expected, got)
			}
		})
	}
}

// BENCHMARKS

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		linkedlist.NewSingly(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}

func BenchmarkBeginInsert(b *testing.B) {
	list := linkedlist.NewSingly()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.BeginInsert(1)
	}
}

func BenchmarkLastInsert(b *testing.B) {
	list := linkedlist.NewSingly()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.LastInsert(1)
	}
}

func BenchmarkBeginDelete(b *testing.B) {
	list := linkedlist.NewSingly(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.BeginDelete()
	}
}

func BenchmarkLastDelete(b *testing.B) {
	list := linkedlist.NewSingly(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		list.LastDelete()
	}
}
