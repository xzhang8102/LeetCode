package golang

/*
 * @lc app=leetcode.cn id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	totalSlots, row, col, block := 0, map[int]map[byte]bool{}, map[int]map[byte]bool{}, map[int]map[byte]bool{}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if row[r] == nil {
				row[r] = map[byte]bool{}
			}
			if col[c] == nil {
				col[c] = map[byte]bool{}
			}
			blockId := (r/3)*3 + c/3
			if block[blockId] == nil {
				block[blockId] = map[byte]bool{}
			}
			char := board[r][c]
			if char != '.' {
				row[r][char] = true
				col[c][char] = true
				block[blockId][char] = true
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
			availble := []byte{}
			for char := byte('9'); char > '0'; char-- {
				if !row[r][char] && !col[c][char] && !block[blockId][char] {
					availble = append(availble, char)
				}
			}
			for _, fill := range availble {
				board[r][c] = fill
				row[r][fill] = true
				col[c][fill] = true
				block[blockId][fill] = true
				filled++
				backtrack(r, c)
				if filled == totalSlots {
					return
				}
				board[r][c] = '.'
				row[r][fill] = false
				col[c][fill] = false
				block[blockId][fill] = false
				filled--
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
