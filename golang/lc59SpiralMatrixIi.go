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
	directions := [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	r, c, dir := 0, 0, 0
	for fill <= n*n {
		ans[r][c] = fill
		fill++
		newR, newC := r+directions[dir][0], c+directions[dir][1]
		if newR < 0 || newR >= n || newC < 0 || newC >= n || ans[newR][newC] != 0 {
			dir = (dir + 1) % 4
		}
		r, c = r+directions[dir][0], c+directions[dir][1]
	}
	return ans
}

// @lc code=end
