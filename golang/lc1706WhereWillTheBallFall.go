package golang

/*
 * @lc app=leetcode.cn id=1706 lang=golang
 *
 * [1706] Where Will the Ball Fall
 */

// @lc code=start
func findBall(grid [][]int) []int {
	row, col := len(grid), len(grid[0])
	ans := make([]int, col)
	var slide func(r, c int) int
	slide = func(r, c int) int {
		if grid[r][c] == 1 {
			if c == col-1 || grid[r][c+1] == -1 {
				return -1
			} else {
				if r == row-1 {
					return c + 1
				}
				return slide(r+1, c+1)
			}
		} else {
			if c == 0 || grid[r][c-1] == 1 {
				return -1
			} else {
				if r == row-1 {
					return c - 1
				}
				return slide(r+1, c-1)
			}
		}
	}
	for c := 0; c < col; c++ {
		ans[c] = slide(0, c)
	}
	return ans
}

// @lc code=end
