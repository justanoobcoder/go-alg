package sort

func DNFSort(a []int, mid int) []int {
	n := len(a)
	i, j, k := 0, 0, n-1
	for j <= k {
		if a[j] < mid {
			a[i], a[j] = a[j], a[i]
			i++
			j++
		} else if a[j] > mid {
			a[j], a[k] = a[k], a[j]
			k--
		} else {
			j++
		}
	}

	return a
}
