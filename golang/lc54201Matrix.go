package golang

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 Matrix
 */

// @lc code=start
func updateMatrix(mat [][]int) [][]int {
	row, col := len(mat), len(mat[0])
	ans := make([][]int, row)
	for i := 0; i < row; i++ {
		ans[i] = make([]int, col)
	}
	queue := make([][2]int, 0)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if mat[r][c] == 0 {
				queue = append(queue, [2]int{r, c})
				// ans[r][c] = 0
				mat[r][c] = '#'
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
		pair := queue[0]
		r, c := pair[0], pair[1]
		for _, dir := range directions {
			newr, newc := r+dir[0], c+dir[1]
			if newr < row && newr >= 0 && newc < col && newc >= 0 && mat[newr][newc] != '#' {
				// enqueue
				queue = append(queue, [2]int{newr, newc})
				// set answer
				ans[newr][newc] = ans[r][c] + 1
				// mark as seen
				mat[newr][newc] = '#'
			}
		}
		queue = queue[1:]
	}
	queue = nil
	return ans
}

// @lc code=end
