package sort

import "slices"

func CountingSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	k := slices.Max(a)
	cnt := make([]int, k+1)
	res := make([]int, n)

	for i := 0; i < n; i++ {
		cnt[a[i]]++
	}

	for i := 1; i <= k; i++ {
		cnt[i] += cnt[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		curr := a[i]
		j := cnt[curr] - 1
		cnt[curr]--
		res[j] = curr
	}

	return res
}
