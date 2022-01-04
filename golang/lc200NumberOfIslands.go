package golang

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	ans := 0
	row := len(grid)
	if row == 0 {
		return ans
	}
	col := len(grid[0])
	queue := [][2]int{}
	directions := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] != '0' {
				ans++
				grid[r][c] = '0'
				queue = append(queue, [2]int{r, c})
				for len(queue) > 0 {
					tmp := queue
					queue = [][2]int{}
					for _, pair := range tmp {
						currR, currC := pair[0], pair[1]
						for _, dir := range directions {
							newR, newC := currR+dir[0], currC+dir[1]
							if newR >= 0 && newR < row && newC >= 0 && newC < col && grid[newR][newC] != '0' {
								grid[newR][newC] = '0'
								queue = append(queue, [2]int{newR, newC})
							}
						}
					}
					tmp = nil
				}
			}
		}
	}
	return ans
}

// @lc code=end
