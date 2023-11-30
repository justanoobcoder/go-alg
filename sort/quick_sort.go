package sort

func partition(arr []int, start, end int) int {
	pivot := arr[end]
	p := start - 1

	for i := start; i < end; i++ {
		if arr[i] <= pivot {
			p++
			arr[i], arr[p] = arr[p], arr[i]
		}
	}
	p++
	arr[p], arr[end] = arr[end], arr[p]

	return p
}

func quickSortWithRange(arr []int, start, end int) {
	if len(arr) <= 1 {
		return
	}
	if start < end {
		p := partition(arr, start, end)
		quickSortWithRange(arr, start, p-1)
		quickSortWithRange(arr, p+1, end)
	}
}

// QuickSort
// Best case: O(nlog(n))
// Worst case: O(n^2)
// Avg: O(nlog(n))
func QuickSort(arr []int) []int {
	if len(arr) > 1 {
		quickSortWithRange(arr, 0, len(arr)-1)
	}
	return arr
}
