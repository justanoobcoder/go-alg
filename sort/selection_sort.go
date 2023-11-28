package sort

// SelectionSort
// Worst case: O(n^2)
// Best case: O(n^2)
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if i != minIdx {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}

	return arr
}
