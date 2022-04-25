package golang

/*
 * @lc app=leetcode.cn id=883 lang=golang
 *
 * [883] Projection Area of 3D Shapes
 */

// @lc code=start
func projectionArea(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	ans := 0
	first := 0
	for i := 0; i < n; i++ {
		first = lc883Max(first, grid[i][0])
		for j := 0; j < m; j++ {
			if grid[i][j] > 0 {
				ans++
			}
			grid[i][0] = lc883Max(grid[i][0], grid[i][j])
			if j > 0 {
				grid[0][j] = lc883Max(grid[0][j], grid[i][j])
			}
		}
		ans += grid[i][0]
	}
	ans += first
	for _, v := range grid[0][1:] {
		ans += v
	}
	return ans
}

func lc883Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
