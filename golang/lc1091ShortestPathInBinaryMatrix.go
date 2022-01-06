package golang

/*
 * @lc app=leetcode.cn id=1091 lang=golang
 *
 * [1091] Shortest Path in Binary Matrix
 */

// @lc code=start

func shortestPathBinaryMatrix(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return -1
	}
	col := len(grid[0])
	if grid[0][0] == 1 || grid[row-1][col-1] == 1 {
		return -1
	}
	// counter-clock wise
	directions := [8][2]int{
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
		{-1, -1},
		{0, -1},
		{1, -1},
	}
	queue := [][2]int{{0, 0}}
	grid[0][0] = 1
	for len(queue) > 0 {
		tmp := queue
		queue = [][2]int{}
		for _, coord := range tmp {
			r, c := coord[0], coord[1]
			for _, dir := range directions {
				newR, newC := coord[0]+dir[0], coord[1]+dir[1]
				if newR >= 0 && newR < row && newC >= 0 && newC < col && grid[newR][newC] == 0 {
					queue = append(queue, [2]int{newR, newC})
					grid[newR][newC] = grid[r][c] + 1
				}
			}
		}
	}
	if grid[row-1][col-1] == 0 {
		return -1
	}
	return grid[row-1][col-1]
}

// @lc code=end
