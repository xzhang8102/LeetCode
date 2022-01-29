package golang

/*
 * @lc app=leetcode.cn id=1765 lang=golang
 *
 * [1765] Map of Highest Peak
 */

// @lc code=start
func highestPeak(isWater [][]int) [][]int {
	row, col := len(isWater), len(isWater[0])
	ans := make([][]int, row)
	for r := range ans {
		ans[r] = make([]int, col)
		for c := range ans[r] {
			ans[r][c] = -1
		}
	}
	type pair struct{ row, col int }
	q := []pair{}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if isWater[r][c] == 1 {
				ans[r][c] = 0
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
			h := ans[r][c]
			for _, dir := range directions {
				newR, newC := r+dir[0], c+dir[1]
				if newR >= 0 && newR < row && newC >= 0 && newC < col && ans[newR][newC] == -1 {
					ans[newR][newC] = h + 1 // visited height must be lower than the height currently visiting
					q = append(q, pair{newR, newC})
				}
			}
		}
	}
	return ans
}

// @lc code=end
