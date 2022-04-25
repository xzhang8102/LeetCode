package golang

/*
 * @lc app=leetcode.cn id=883 lang=golang
 *
 * [883] Projection Area of 3D Shapes
 */

// @lc code=start
func projectionArea(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	row, col := make([]int, n), make([]int, m)
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			row[i] = lc883Max(row[i], grid[i][j])
			col[j] = lc883Max(col[j], grid[i][j])
			if grid[i][j] > 0 {
				ans++
			}
		}
	}
	for _, v := range row {
		ans += v
	}
	for _, v := range col {
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
