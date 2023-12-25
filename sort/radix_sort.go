package sort

import "slices"

func RadixSort(a []int) []int {
	n := len(a)
	if n <= 1 {
		return a
	}
	max := slices.Max(a)
	for i := 1; max/i > 0; i *= 10 {
		a = countingSort(a, n, i)
	}
	return a
}

func countingSort(a []int, n, pos int) []int {
	k := 10
	cnt := make([]int, k)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		j := (a[i] / pos) % 10
		cnt[j]++
	}

	for i := 1; i < k; i++ {
		cnt[i] += cnt[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		j := (a[i] / pos) % 10
		cnt[j]--
		res[cnt[j]] = a[i]
	}

	return res
}
