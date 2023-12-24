package sort

// QuickSort
// Best case: O(nlog(n))
// Worst case: O(n^2) (when dataset's already sorted,
//    but can be avoided and reduced to O(nlog(n)) )
// Avg: O(nlog(n))
func QuickSort(arr []int) []int {
	quickSortWithRange(arr, 0, len(arr)-1)
	return arr
}

func partition(arr []int, start, end int) int {
	mid := (start + end) / 2
	if arr[mid] < arr[start] {
		arr[mid], arr[start] = arr[start], arr[mid]
	}
	if arr[end] < arr[start] {
		arr[end], arr[start] = arr[start], arr[end]
	}
	if arr[mid] < arr[end] {
		arr[mid], arr[end] = arr[end], arr[mid]
	}
	pivot := arr[end]
	i := start - 1
	j := end + 1
	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}
		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func quickSortWithRange(arr []int, start, end int) {
	if start < end {
		p := partition(arr, start, end)
		quickSortWithRange(arr, start, p)
		quickSortWithRange(arr, p+1, end)
	}
}
