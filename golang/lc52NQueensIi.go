package golang

/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start
func totalNQueens(n int) int {
	col, left, right := map[int]int{}, map[int]int{}, map[int]int{}
	ans := 0
	var backtrack func(r int)
	backtrack = func(r int) {
		if r == n {
			ans++
			return
		}
		for c := 0; c < n; c++ {
			if col[c] == 0 && left[r-c] == 0 && right[r+c] == 0 {
				col[c]++
				left[r-c]++
				right[r+c]++
				backtrack(r + 1)
				col[c]--
				left[r-c]--
				right[r+c]--
			}
		}
	}
	backtrack(0)
	return ans
}

// @lc code=end
