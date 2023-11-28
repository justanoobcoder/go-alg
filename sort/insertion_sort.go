package sort

// InsertionSort
// Worst case: O(n^2)
// Best case: O(n)
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		temp := arr[i]
		idx := i
		for ; idx > 0 && arr[idx-1] > temp; idx-- {
			arr[idx] = arr[idx-1]
		}
		arr[idx] = temp
	}

	return arr
}
