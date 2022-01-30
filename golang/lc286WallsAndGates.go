package golang

import "math"

/*
 * @lc app=leetcode.cn id=286 lang=golang
 *
 * [286] Walls and Gates
 */

// @lc code=start
func wallsAndGates(rooms [][]int) {
	row := len(rooms)
	if row == 0 {
		return
	}
	col := len(rooms[0])
	type pair struct{ row, col int }
	q := []pair{}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if rooms[r][c] == 0 {
				q = append(q, pair{r, c})
			}
		}
	}
	directions := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	for len(q) > 0 {
		tmp := q
		q = []pair{}
		for _, p := range tmp {
			r, c := p.row, p.col
			d := rooms[r][c]
			for _, dir := range directions {
				newR, newC := r+dir[0], c+dir[1]
				if newR >= 0 && newR < row && newC >= 0 && newC < col && rooms[newR][newC] == math.MaxInt32 {
					rooms[newR][newC] = d + 1
					q = append(q, pair{newR, newC})
				}
			}
		}
	}
}

func lc286Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
