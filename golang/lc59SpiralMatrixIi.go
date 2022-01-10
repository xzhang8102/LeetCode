package golang

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		ans[i] = make([]int, n)
	}
	fill := 1
	for counter := 0; n > 0; counter, n = counter+1, n-1 {
		row, col := counter, counter
		for ; col < n; col++ {
			ans[row][col] = fill
			fill++
		}
		col--
		row++
		for ; row < n; row++ {
			ans[row][col] = fill
			fill++
		}
		row--
		col--
		for ; col >= counter; col-- {
			ans[row][col] = fill
			fill++
		}
		col++
		row--
		for ; row > counter; row-- {
			ans[row][col] = fill
			fill++
		}
	}
	return ans
}

// @lc code=end
