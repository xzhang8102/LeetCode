package golang

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	totalSlots, row, col, block := 0, [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			char := board[r][c]
			if char != '.' {
				blockId := (r/3)*3 + c/3
				row[r][char-'1'] = true
				col[c][char-'1'] = true
				block[blockId][char-'1'] = true
			} else {
				totalSlots++
			}
		}
	}
	var backtrack func(int, int)
	filled := 0
	backtrack = func(r, c int) {
		if board[r][c] == '.' {
			blockId := (r/3)*3 + c/3
			for char := byte('9'); char > '0'; char-- {
				if !row[r][char-'1'] && !col[c][char-'1'] && !block[blockId][char-'1'] {
					board[r][c] = char
					row[r][char-'1'] = true
					col[c][char-'1'] = true
					block[blockId][char-'1'] = true
					filled++
					backtrack(r, c)
					if filled == totalSlots {
						return
					}
					board[r][c] = '.'
					row[r][char-'1'] = false
					col[c][char-'1'] = false
					block[blockId][char-'1'] = false
					filled--
				}
			}
		} else {
			if filled == totalSlots {
				return
			}
			newR, newC := r, c+1
			if newC > 8 {
				newR++
				newC = 0
			}
			backtrack(newR, newC)
		}
	}
	backtrack(0, 0)
}

// @lc code=end
