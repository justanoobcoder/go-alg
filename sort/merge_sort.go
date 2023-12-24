package sort

// MergeSort
// Best case: O(nlog(n))
// Worst case: O(nlog(n))
// Avg case: O(nlog(n))
func MergeSort(arr []int) []int {
	recursiveMergeSort(arr, 0, len(arr)-1)
	return arr
}

func merge(a []int, start, mid, end int) {
	n1 := mid - start + 1
	n2 := end - mid
	a1 := make([]int, n1)
	a2 := make([]int, n2)
	copy(a1, a[start:(start+n1)])
	copy(a2, a[(mid+1):(mid+1+n2)])
	i, j, k := 0, 0, start
	for i < n1 && j < n2 {
		if a1[i] < a2[j] {
			a[k] = a1[i]
			i++
		} else {
			a[k] = a2[j]
			j++
		}
		k++
	}
	for i < n1 {
		a[k] = a1[i]
		i++
		k++
	}
	for j < n2 {
		a[k] = a2[j]
		j++
		k++
	}
}

func recursiveMergeSort(a []int, start, end int) {
	if start < end {
		mid := (start + end) / 2
		recursiveMergeSort(a, start, mid)
		recursiveMergeSort(a, mid+1, end)
		merge(a, start, mid, end)
	}
}
