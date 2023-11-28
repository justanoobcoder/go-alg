package sort_test

import (
	"github.com/justanoobcoder/go-alg/sort"
	"github.com/stretchr/testify/assert"
	"testing"
)

func runTest(t *testing.T, sortFunc func([]int) []int) {
	tests := getTests()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			want := test.expected
			got := sortFunc(test.input)
			assert.ElementsMatch(t, want, got)
		})
	}
}

func runBenchmark(b *testing.B, sortFunc func([]int) []int) {
	tests := getTests()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			sortFunc(test.input)
		}
	}
}

// TESTS

func TestBubbleSort(t *testing.T) {
	runTest(t, sort.BubbleSort)
}

// BENCHMARKS

func BenchmarkBubbleSort(b *testing.B) {
	runBenchmark(b, sort.BubbleSort)
}

type Test struct {
	name     string
	input    []int
	expected []int
}

func getTests() []Test {
	return []Test{
		{
			name:     "Sorted Unsigned",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Reversed Unsigned",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Sorted Signed",
			input:    []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Reversed Signed",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Reversed Signed #2",
			input:    []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, -1, -2, -3, -4, -5, -6, -7, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Random order Signed",
			input:    []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -10},
			expected: []int{-10, -9, -8, -7, -6, -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 10},
		},
		{
			name:     "Singleton",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "Empty Slice",
			input:    []int{},
			expected: []int{},
		},
	}
}