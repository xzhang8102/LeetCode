package golang

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	slots, row, col, block := [][]int{}, [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			char := board[r][c]
			if char != '.' {
				blockId := (r/3)*3 + c/3
				row[r][char-'1'] = true
				col[c][char-'1'] = true
				block[blockId][char-'1'] = true
			} else {
				slots = append(slots, []int{r, c})
			}
		}
	}
	var backtrack func(int) bool
	backtrack = func(pos int) bool {
		if pos == len(slots) {
			return true
		}
		r, c := slots[pos][0], slots[pos][1]
		blockId := (r/3)*3 + c/3
		for char := byte('9'); char > '0'; char-- {
			if !row[r][char-'1'] && !col[c][char-'1'] && !block[blockId][char-'1'] {
				board[r][c] = char
				row[r][char-'1'] = true
				col[c][char-'1'] = true
				block[blockId][char-'1'] = true
				if backtrack(pos + 1) {
					return true
				}
				board[r][c] = '.'
				row[r][char-'1'] = false
				col[c][char-'1'] = false
				block[blockId][char-'1'] = false
			}
		}
		return false
	}
	backtrack(0)
}

// @lc code=end
