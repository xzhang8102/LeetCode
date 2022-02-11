package golang

/*
 * @lc app=leetcode.cn id=52 lang=golang
 *
 * [52] N-Queens II
 */

// @lc code=start
func totalNQueens(n int) int {
	col, left, right := map[int]int{}, map[int]int{}, map[int]int{}
	ans := 0
	pick := make([]int, 0, n)
	var backtrack func(r, c int)
	backtrack = func(r, c int) {
		if len(pick) == n-1 {
			ans++
			return
		}
		pick = append(pick, c)
		col[c]++
		left[r-c]++
		right[r+c]++
		for newR, newC := r+1, 0; newC < n && newR < n; newC++ {
			if col[newC] == 0 && left[newR-newC] == 0 && right[newR+newC] == 0 {
				backtrack(newR, newC)
			}
		}
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
