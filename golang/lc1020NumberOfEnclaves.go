package golang

/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] Number of Enclaves
 */

// @lc code=start
func numEnclaves(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	col := len(grid[0])
	q := [][]int{}
	for c := 0; c < col; c++ {
		if grid[0][c] == 1 {
			q = append(q, []int{0, c})
		}
		if grid[row-1][c] == 1 {
			q = append(q, []int{row - 1, c})
		}
	}
	for r := 1; r < row-1; r++ {
		if grid[r][0] == 1 {
			q = append(q, []int{r, 0})
		}
		if grid[r][col-1] == 1 {
			q = append(q, []int{r, col - 1})
		}
	}
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		r, c := curr[0], curr[1]
		// chances that elements during initialization
		// have already been visited
		if grid[r][c] == 0 {
			continue
		}
		grid[r][c] = 0
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < row && newC >= 0 && newC < col && grid[newR][newC] == 1 {
				q = append(q, []int{newR, newC})
			}
		}
	}
	ans := 0
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 1 {
				ans++
			}
		}
	}
	return ans
}

// @lc code=end
