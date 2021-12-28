package golang

/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] Max Area of Island
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	direction := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	ans := 0
	row, col := len(grid), len(grid[0])
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		grid[r][c] = 0
		area := 1
		for _, dir := range direction {
			newr, newc := r+dir[0], c+dir[1]
			if newr >= 0 && newr < row && newc >= 0 && newc < col && grid[newr][newc] != 0 {
				area += dfs(newr, newc)
			}
		}
		return area
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 1 {
				area := dfs(r, c)
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}

// @lc code=end
