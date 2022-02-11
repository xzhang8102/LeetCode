package golang

/*
 * @lc app=leetcode.cn id=51 lang=golang
 *
 * [51] N-Queens
 */

// @lc code=start
func solveNQueens(n int) [][]string {
	row, col, left, right := map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	ans := [][]string{}
	pick := make([]int, 0, n)
	var backtrack func(r, c int)
	backtrack = func(r, c int) {
		pick = append(pick, c)
		row[r]++
		col[c]++
		left[r-c]++
		right[r+c]++
		for newR, newC := r+1, 0; newC < n && newR < n; newC++ {
			if row[newR] == 0 && col[newC] == 0 && left[newR-newC] == 0 && right[newR+newC] == 0 {
				backtrack(newR, newC)
			}
		}
		if len(pick) == n {
			tmp := make([]byte, n)
			for i := range tmp {
				tmp[i] = '.'
			}
			str := []string{}
			for _, pos := range pick {
				tmp[pos] = 'Q'
				str = append(str, string(tmp))
				tmp[pos] = '.'
			}
			ans = append(ans, str)
		}
		row[r]--
		col[c]--
		left[r-c]--
		right[r+c]--
		pick = pick[:len(pick)-1]
	}
	for i := 0; i < n; i++ {
		backtrack(0, i)
	}
	return ans
}

// @lc code=end
