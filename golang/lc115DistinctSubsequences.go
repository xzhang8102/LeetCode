package golang

/*
 * @lc app=leetcode.cn id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */

// @lc code=start
func numDistinct(s string, t string) int {
	n := len(t)
	ans := 0
	var backtrack func(si, ti int)
	backtrack = func(si, ti int) {
		if ti == n {
			ans++
			return
		}
		for i := si; i < len(s); i++ {
			if s[i] == t[ti] {
				backtrack(i+1, ti+1)
			}
		}
	}
	backtrack(0, 0)
	return ans
}

// @lc code=end
