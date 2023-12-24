package sort

// InsertionSort
// Worst case: O(n^2)
// Best case: O(n) (when dataset's already sorted)
// Avg case: O(n^2)
func InsertionSort(arr []int) []int {
	n := len(arr)
	for i := 1; i < n; i++ {
		curr := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > curr; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = curr
	}

	return arr
}
