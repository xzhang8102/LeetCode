package golang

/*
 * @lc app=leetcode.cn id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start
func pacificAtlantic(heights [][]int) [][]int {
	row := len(heights)
	if row == 0 {
		return nil
	}
	col := len(heights[0])
	pacific, atlantic := make([][]int, row), make([][]int, row)
	for i := range pacific {
		pacific[i] = make([]int, col)
		atlantic[i] = make([]int, col)
	}
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(int, int, [][]int)
	dfs = func(r, c int, grid [][]int) {
		grid[r][c] = 1
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < row && newC >= 0 && newC < col {
				if grid[newR][newC] != 1 && heights[newR][newC] >= heights[r][c] {
					dfs(newR, newC, grid)
				}
			}
		}
	}
	for i := 0; i < row; i++ {
		dfs(i, 0, pacific)
		dfs(i, col-1, atlantic)
	}
	for i := 0; i < col; i++ {
		dfs(0, i, pacific)
		dfs(row-1, i, atlantic)
	}
	ans := [][]int{}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if pacific[i][j] == 1 && atlantic[i][j] == 1 {
				ans = append(ans, []int{i, j})
			}
		}
	}
	return ans
}

// @lc code=end
