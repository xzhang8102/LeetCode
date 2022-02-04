package golang

/*
 * @lc app=leetcode.cn id=1219 lang=golang
 *
 * [1219] Path with Maximum Gold
 */

// @lc code=start
func getMaximumGold(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	ans := 0
	pick := 0
	directions := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	var backtrack func(int, int)
	backtrack = func(r, c int) {
		val := grid[r][c]
		pick += val
		grid[r][c] = 0
		if pick > ans {
			ans = pick
		}
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < row && newC >= 0 && newC < col && grid[newR][newC] != 0 {
				backtrack(newR, newC)
			}
		}
		pick -= val
		grid[r][c] = val
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] != 0 {
				backtrack(r, c)
			}
		}
	}
	return ans
}

// @lc code=end
