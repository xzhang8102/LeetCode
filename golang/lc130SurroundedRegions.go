package golang

/*
 * @lc app=leetcode.cn id=130 lang=golang
 *
 * [130] Surrounded Regions
 */

// @lc code=start
func solve(board [][]byte) {
	row := len(board)
	if row == 0 {
		return
	}
	col := len(board[0])
	directions := [4][2]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= row || c < 0 || c >= col || board[r][c] != 'O' {
			return
		}
		board[r][c] = 'A'
		for _, dir := range directions {
			dfs(r+dir[0], c+dir[1])
		}
	}
	for r := 0; r < row; r++ {
		dfs(r, 0)
		dfs(r, col-1)
	}
	for c := 1; c < col-1; c++ {
		dfs(0, c)
		dfs(row-1, c)
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if board[r][c] == 'A' {
				board[r][c] = 'O'
			} else if board[r][c] == 'O' {
				board[r][c] = 'X'
			}
		}
	}
}

// @lc code=end
