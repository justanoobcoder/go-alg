package sort

// BubbleSort
// Worst case: O(n^2)
// Best case: O(n)
func BubbleSort(arr []int) []int {
	swapped := true
	var sortedElems int
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1-sortedElems; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				swapped = true
			}
		}
		sortedElems++
	}
	return arr
}
