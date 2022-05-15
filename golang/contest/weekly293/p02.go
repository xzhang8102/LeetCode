package weekly293

import "sort"

func maxConsecutive(bottom int, top int, special []int) int {
	sort.Ints(special)
	n := len(special)
	ans := special[0] - bottom
	for i := 1; i < n; i++ {
		ans = max(ans, special[i]-special[i-1]-1)
	}
	ans = max(ans, top-special[n-1])
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
