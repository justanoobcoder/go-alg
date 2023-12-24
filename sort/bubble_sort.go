package sort

// BubbleSort
// Worst case: O(n^2)
// Best case: O(n) (when dataset's already sorted)
// Avg case: O(n^2)
func BubbleSort(arr []int) []int {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n--
	}

	return arr
}
