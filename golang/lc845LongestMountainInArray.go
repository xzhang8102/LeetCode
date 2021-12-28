package golang

/*
 * @lc app=leetcode.cn id=845 lang=golang
 *
 * [845] Longest Mountain in Array
 */

// @lc code=start
func longestMountain(arr []int) int {
	if len(arr) < 3 {
		return 0
	}
	n := len(arr)
	// (left|right)[i]: the number of elements monotonically decreasing from i
	left, right := make([]int, n), make([]int, n)
	for i := 1; i < n-1; i++ {
		if arr[i] > arr[i-1] {
			left[i] = left[i-1] + 1
		}
	}
	for i := n - 2; i > 0; i-- {
		if arr[i] > arr[i+1] {
			right[i] = right[i+1] + 1
		}
	}
	ans := 0
	for i, l := range left {
		r := right[i]
		if l > 0 && r > 0 && l+r+1 > ans {
			ans = l + r + 1
		}
	}
	return ans
}

// @lc code=end
