package sort

// CountingSort
// Time complexity: O(n+k)
// Space complexity: O(n+k)
// Stable: Yes
func CountingSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	ma, mi := a[0], a[0]
	for _, num := range a {
		ma = max(ma, num)
		mi = min(mi, num)
	}
	cnt := make([]int, ma-mi+1)
	res := make([]int, n)

	for _, num := range a {
		cnt[num-mi]++
	}

	for i := 1; i < len(cnt); i++ {
		cnt[i] += cnt[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		curr := a[i] - mi
		j := cnt[curr] - 1
		cnt[curr]--
		res[j] = a[i]
	}

	return res
}
