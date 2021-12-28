package golang

/*
 * @lc app=leetcode.cn id=994 lang=golang
 *
 * [994] Rotting Oranges
 */

// @lc code=start
func orangesRotting(grid [][]int) int {
	row, col := len(grid), len(grid[0])
	ans := 0
	fresh := 0
	queue := make([][2]int, 0)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, [2]int{r, c})
			}
			if grid[r][c] == 1 {
				fresh++
			}
		}
	}
	directions := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for len(queue) > 0 {
		size := len(queue)
		affected := false
		for _, pair := range queue {
			r, c := pair[0], pair[1]
			for _, dir := range directions {
				newr, newc := r+dir[0], c+dir[1]
				if newr >= 0 && newr < row && newc >= 0 && newc < col && grid[newr][newc] == 1 {
					affected = true
					fresh--
					grid[newr][newc] = 2
					queue = append(queue, [2]int{newr, newc})
				}
			}
		}
		if affected {
			ans++
		}
		queue = queue[size:]
	}
	queue = nil
	if fresh > 0 {
		return -1
	}
	return ans
}

// @lc code=end
