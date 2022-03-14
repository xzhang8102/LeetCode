package golang

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	row, col, blk := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < 9; i++ {
		row[i] = make([]bool, 10)
		col[i] = make([]bool, 10)
		blk[i] = make([]bool, 10)
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] != '.' {
				num := board[r][c] - '0'
				blkId := (r/3)*3 + c/3
				if row[r][num] || col[c][num] || blk[blkId][num] {
					return false
				} else {
					row[r][num] = true
					col[c][num] = true
					blk[blkId][num] = true
				}
			}
		}
	}
	return true
}

// @lc code=end
