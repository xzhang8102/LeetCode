package golang

/*
 * @lc app=leetcode.cn id=2055 lang=golang
 *
 * [2055] Plates Between Candles
 */

// @lc code=start
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	// assume the left end is '|'
	sum := make([]int, n+1)
	// index of the left closest '|'; index of the right closest '|'
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == '*' {
			sum[i+1] = sum[i] + 1
			if i == 0 {
				left[i] = -1
			} else {
				left[i] = left[i-1]
			}
		} else {
			sum[i+1] = sum[i]
			left[i] = i
		}
	}
	for i := n - 1; i >= 0; i-- {
		if s[i] == '*' {
			if i == n-1 {
				right[i] = -1
			} else {
				right[i] = right[i+1]
			}
		} else {
			right[i] = i
		}
	}
	ans := []int{}
	for _, query := range queries {
		l, r := right[query[0]], left[query[1]]
		if l <= r && l >= query[0] && r <= query[1] {
			ans = append(ans, sum[r+1]-sum[l+1])
		} else {
			ans = append(ans, 0)
		}
	}
	return ans
}

// @lc code=end
