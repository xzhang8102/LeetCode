package golang

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0 || word == "" {
		return false
	}
	col := len(board[0])
	directions := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	n := len(word)
	var dfs func(int, int, int) bool
	dfs = func(r, c, index int) bool {
		if board[r][c] != word[index] {
			return false
		}
		if index == n-1 {
			return true
		}
		char := board[r][c]
		board[r][c] = '#'
		defer func() {
			board[r][c] = char
		}()
		for _, dir := range directions {
			newR, newC := r+dir[0], c+dir[1]
			if newR >= 0 && newR < row && newC >= 0 && newC < col && board[newR][newC] != '#' {
				if dfs(newR, newC, index+1) {
					return true
				}
			}
		}
		return false
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
